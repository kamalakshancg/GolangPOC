package com.poc.entity;

import lombok.Getter;

import java.util.HashSet;
import java.util.Set;

@Getter
public class User {
    private final Integer id;
    private final String name;
    private final String email;
    private final Set<Order> orders = new HashSet<>();

    public User(Integer id, String name, String email) {
        this.id = id;
        this.name = name;
        this.email = email;
    }

    public void addOrder(final Order order){
        this.orders.add(order);
    }
}