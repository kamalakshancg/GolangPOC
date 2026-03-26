package com.poc.entity;

import lombok.Getter;

import java.util.HashSet;
import java.util.Set;

@Getter
public class Order {
    private final Integer id;
    private final Double amount;
    private final String status;
    private final String description;
    private final Integer userId;
    private Set<Item> items = new HashSet<>();

    public Order(Integer id, Double amount, String status, String description, Integer userId) {
        this.id = id;
        this.amount = amount;
        this.status = status;
        this.description = description;
        this.userId = userId;
    }

    public void addItem(final Item item){
        this.items.add(item);
    }
}