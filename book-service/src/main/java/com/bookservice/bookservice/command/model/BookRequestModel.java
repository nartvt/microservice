package com.bookservice.bookservice.command.model;

import lombok.*;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class BookRequestModel {
    private String bookId;
    private String name;
    private String author;
    private boolean isReady;
}
