package com.bookservice.bookservice.command.event;

import com.bookservice.bookservice.command.data.Book;
import com.bookservice.bookservice.command.data.IBookRepository;
import org.axonframework.eventhandling.EventHandler;
import org.springframework.beans.BeanUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

@Component
public class BookEventHandler {

    @Autowired
    private IBookRepository bookRepository;

    @EventHandler
    public void on(CreateBookEvent event) {
        Book book = new Book();
        BeanUtils.copyProperties(event, book);
        bookRepository.save(book);
    }

    @EventHandler
    public void on(UpdateBookEvent event) {
        final Book book = bookRepository.getReferenceById(event.getBookId());
        book.setAuthor(event.getAuthor());
        book.setName(event.getName());
        book.setReady(event.isReady());
        bookRepository.save(book);
    }

    @EventHandler
    public void on(DeleteBookEvent event) {
        final Book book = bookRepository.getReferenceById(event.getBookId());
        if (book == null) {
            return;
        }
        bookRepository.deleteById(book.getBookId());
    }
}
