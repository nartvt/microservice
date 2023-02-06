package com.nartvt.employeeservice.command.data;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "employees")
@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class Employee {
    @Id
    private String employeeId;
    private String firstName;
    private String lastName;
    private String kin;
    private Boolean isDisciplined;
}
