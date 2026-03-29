package com.poc.entity;

import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;
import jakarta.persistence.Transient;

import java.util.HashSet;
import java.util.Set;


@Entity
@Table(name = "users")
public class User {
    @Id
    private Integer id;
    private String name;
    private String email;
    @Transient
    private final Set<Order> orders = new HashSet<>();

    public User(Integer id, String name, String email) {
        this.id = id;
        this.name = name;
        this.email = email;
    }

    protected User(){
    }

    public Integer getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public String getEmail() {
        return email;
    }

    public Set<Order> getOrders() {
        return orders;
    }

    public void addOrder(final Order order){
        this.orders.add(order);
    }
}