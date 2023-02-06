package com.bookservice.bookservice.command.event;

import javax.persistence.EntityNotFoundException;

import org.axonframework.eventhandling.EventHandler;
import org.springframework.beans.BeanUtils;
import org.springframework.stereotype.Component;

import com.bookservice.bookservice.command.data.Book;
import com.bookservice.bookservice.command.data.IBookRepository;

@Component
public class BookEventHandler {
	private final IBookRepository bookRepository;

	public BookEventHandler(final IBookRepository bookRepository) {
		this.bookRepository = bookRepository;
	}

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
		if (event == null || event.getBookId() == null) {
			return;
		}
		try {
			bookRepository.deleteById(event.getBookId());
		} catch (IllegalArgumentException e) {
			System.out.println(e.getMessage());
		}
	}
}
