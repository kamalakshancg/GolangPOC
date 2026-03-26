package com.poc.dto;

import com.poc.entity.Item;

import java.util.ArrayList;
import java.util.List;

public class NestedOrder {
    public int id;
    public double amount;
    public List<Item> items = new ArrayList<>();

    public NestedOrder(int id, double amount, List<Item> items) {
        this.id = id;
        this.amount = amount;
        this.items = items;
    }
}