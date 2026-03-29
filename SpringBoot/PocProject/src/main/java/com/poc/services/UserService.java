package com.poc.services;

import com.poc.entity.Item;
import com.poc.entity.Order;
import com.poc.entity.User;
import com.poc.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Service
public class UserService {

    @Autowired
    private UserRepository repo;

    public List<User> getUserWithOrderDetails() {
        List<Object[]> rows = repo.getUserWithOrders();
        // Map to group Users by ID while maintaining SQL Order
        Map<Integer, User> userMap = new LinkedHashMap<>();
        for (Object[] row : rows) {
            final int userId = ((Number) row[0]).intValue();
            final String userName = (String) row[1];
            final int orderId = ((Number) row[2]).intValue();
            final double amount = ((Number) row[3]).doubleValue();
            final int itemId = ((Number) row[4]).intValue();
            final String product = (String) row[5];
            final int qty = ((Number) row[6]).intValue();
            final double price = ((Number) row[7]).doubleValue();
            final String email = (String) row[8];
            final String status = (String) row[9];
            final String description = (String) row[10];

            // 1. Get or Create User
            final User user = userMap.computeIfAbsent(userId, id -> new User(id, userName, email));

            // 2. Find or Create Order
            Order currentOrder = null;
            for (Order order : user.getOrders()) {
                if (order.getId() == orderId) {
                    currentOrder = order;
                    break;
                }
            }

            if (currentOrder == null) {
                currentOrder = new Order(orderId, amount, status, description, userId);
                user.addOrder(currentOrder);
            }

            //3. Create and Add Item
            final Item item = new Item(itemId, product, qty, price);
            currentOrder.addItem(item);
        }
        return new ArrayList<>(userMap.values());
    }
}