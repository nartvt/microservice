package com.nartvt.employeeservice.command.model.response;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class EmployeeResponseModel {
    private Object data;
    private int status;
    private String message;
}
