package com.bookservice.bookservice.command.event;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
public class CreateBookEvent {
    private String bookId;
    private String name;
    private String author;
    private boolean isReady;
}
