package com.poc.services;

import com.poc.entity.Order;
import com.poc.repository.OrderRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class OrderService {

    @Autowired
    private OrderRepository orderRepository;

    public List<Order> getOrderDetails(){
        return orderRepository.findWideOrders();
    }
}
