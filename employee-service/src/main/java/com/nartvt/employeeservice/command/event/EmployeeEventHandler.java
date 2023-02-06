package com.nartvt.employeeservice.command.event;

import org.axonframework.eventhandling.EventHandler;
import org.springframework.beans.BeanUtils;
import org.springframework.stereotype.Component;

import com.nartvt.employeeservice.command.data.Employee;
import com.nartvt.employeeservice.command.data.IEmployeeRepository;

@Component
public class EmployeeEventHandler {

	private final IEmployeeRepository employeeRepository;

	public EmployeeEventHandler(IEmployeeRepository employeeRepository) {
		this.employeeRepository = employeeRepository;
	}

	@EventHandler
	public void on(CreateEmployeeEvent event) {
		Employee employee = new Employee();
		BeanUtils.copyProperties(event, employee);
		employeeRepository.save(employee);
	}
}
