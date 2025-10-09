package dependencies

import (
	"log"

	repository "github.com/book-service/api/app/db/repository/employees"
	ServiceEmployees "github.com/book-service/api/app/service/employees"
)

type EmployeeContainer struct {
	employeeService    ServiceEmployees.EmployeeServiceInterface
	employeeRepository repository.EmployeeRepository
}

func (c *Container) GetEmployeeRepository() repository.EmployeeRepository {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.EmployeeContainer == nil {
		c.EmployeeContainer = &EmployeeContainer{}
	}

	if c.EmployeeContainer.employeeRepository == nil {
		c.EmployeeContainer.employeeRepository = repository.NewEmployeeRepository(c.db)
	}
	return c.EmployeeContainer.employeeRepository
}

func (c *Container) GetEmployeeService() ServiceEmployees.EmployeeServiceInterface {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.EmployeeContainer == nil {
		c.EmployeeContainer = &EmployeeContainer{}
	}

	if c.EmployeeContainer.employeeService == nil {
		log.Println("🔧 Getting employee repository...")
		// Don't call GetEmployeeRepository() here - it would cause a deadlock
		// Instead, create the repository directly
		if c.EmployeeContainer.employeeRepository == nil {
			c.EmployeeContainer.employeeRepository = repository.NewEmployeeRepository(c.db)
		}
		employeeRepo := c.EmployeeContainer.employeeRepository
		log.Println("✅ Employee repository obtained")

		log.Println("🔧 Creating employee service...")
		c.EmployeeContainer.employeeService = ServiceEmployees.NewEmployeeService(employeeRepo)
		log.Println("✅ Employee service created")
	}
	return c.EmployeeContainer.employeeService
}
