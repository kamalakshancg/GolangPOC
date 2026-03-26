package com.poc.dto;

import java.util.ArrayList;
import java.util.List;

public class UserWithOrders {
    public int id;
    public String name;
    public List<NestedOrder> orders = new ArrayList<>();


    public UserWithOrders(int id, String name, List<NestedOrder> orders) {
        this.id = id;
        this.name = name;
        this.orders = orders;
    }
}
