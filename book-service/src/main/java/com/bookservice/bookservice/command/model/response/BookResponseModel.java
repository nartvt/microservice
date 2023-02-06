package com.bookservice.bookservice.command.model.response;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
@AllArgsConstructor
public class BookResponseModel {
	private Object data;
	private int status;
	private String message;
}
