package com.nartvt.employeeservice.command.controller;

import com.nartvt.employeeservice.command.command.CreateEmployeeCommand;
import com.nartvt.employeeservice.command.model.request.EmployeeRequestModel;
import com.nartvt.employeeservice.command.model.response.EmployeeResponseModel;
import org.axonframework.commandhandling.CommandExecutionException;
import org.axonframework.commandhandling.distributed.CommandDispatchException;
import org.axonframework.commandhandling.gateway.CommandGateway;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;

@RestController
@RequestMapping("/api/v1/employees")
public class EmployeeCommandController {
    private CommandGateway commandGateway;

    public EmployeeCommandController(CommandGateway commandGateway) {
        this.commandGateway = commandGateway;
    }

    @PostMapping
    public ResponseEntity<EmployeeResponseModel> addEmployee(@RequestBody EmployeeRequestModel model) {
        final CreateEmployeeCommand command = new CreateEmployeeCommand();
        command.setEmployeeId(UUID.randomUUID().toString());
        command.setFirstName(model.getFirstName());
        command.setLastName(model.getLastName());
        command.setKin(model.getKin());
        command.setIsDisciplined(model.getIsDisciplined());

        final EmployeeResponseModel responseModel = new EmployeeResponseModel();
        try {
            commandGateway.sendAndWait(command);
            responseModel.setData(command);
            responseModel.setStatus(HttpStatus.CREATED.value());
            responseModel.setMessage("Success!");
        } catch (CommandExecutionException | CommandDispatchException e) {
            responseModel.setData(null);
            responseModel.setStatus(HttpStatus.INTERNAL_SERVER_ERROR.value());
            responseModel.setMessage(e.getMessage());
        }
        return new ResponseEntity<>(responseModel, HttpStatus.valueOf(responseModel.getStatus()));
    }
}
