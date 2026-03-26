package com.poc.entity;

import lombok.Getter;

@Getter
public class Item {
    private final Integer id;
    private final String product_name;
    private final Integer quantity;
    private final Double price;

    public Item(Integer id, String product_name, Integer quantity, Double price) {
        this.id = id;
        this.product_name = product_name;
        this.quantity = quantity;
        this.price = price;
    }

}