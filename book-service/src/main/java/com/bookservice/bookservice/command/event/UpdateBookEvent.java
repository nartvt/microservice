package com.bookservice.bookservice.command.event;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class UpdateBookEvent {
	private String bookId;
	private String name;
	private String author;
	private boolean isReady;
}