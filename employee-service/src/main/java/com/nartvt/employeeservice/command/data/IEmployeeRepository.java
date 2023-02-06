package com.nartvt.employeeservice.command.data;

import org.springframework.data.jpa.repository.JpaRepository;

public interface IEmployeeRepository extends JpaRepository<Employee, String> {
}
