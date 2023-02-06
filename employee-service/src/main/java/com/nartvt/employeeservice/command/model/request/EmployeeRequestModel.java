package com.nartvt.employeeservice.command.model.request;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class EmployeeRequestModel {
    private String employeeId;
    private String firstName;
    private String lastName;
    private String kin;
    private Boolean isDisciplined;
}