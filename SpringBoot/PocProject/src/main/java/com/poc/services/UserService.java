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

    public List<User> userWithOrder() {
        // 1. Start Stopwatch
        long start = System.currentTimeMillis();

        List<Object[]> rows = repo.userWithOrder();

        // 2. Capture DB Time
        long dbTime = System.currentTimeMillis() - start;
        System.out.println("Test 3 DB Fetch Time: " + dbTime + " ms");

        // Map to group Users by ID while maintaining SQL Order
        Map<Integer, User> userMap = new LinkedHashMap<>();

        for (Object[] row : rows) {
            // Safe Casting for PostgreSQL native query results
            int userId = ((Number) row[0]).intValue();
            String userName = (String) row[1];
            int orderId = ((Number) row[2]).intValue();
            double amount = ((Number) row[3]).doubleValue();
            int itemId = ((Number) row[4]).intValue();
            String product = (String) row[5];
            int qty = ((Number) row[6]).intValue();
            double price = ((Number) row[7]).doubleValue();
            String email = (String) row[8];
            String status = (String) row[9];
            String description = (String) row[10];

            // 1. Get or Create User (Passing null for email since it isn't in our SQL)
            User user = userMap.computeIfAbsent(userId, id -> new User(id, userName, email));

            // 2. Find or Create Order
            Order currentOrder = null;
            // Iterate through existing orders to see if we already created it
            for (Order o : user.getOrders()) {
                if (o.getId() == orderId) {
                    currentOrder = o;
                    break;
                }
            }

            if (currentOrder == null) {
                // Passing null for status/description since they aren't in this SQL projection
                currentOrder = new Order(orderId, amount, status, description, userId);
                user.addOrder(currentOrder);
            }

            // 3. Create and Add Item
            Item item = new Item(itemId, product, qty, price);
            currentOrder.addItem(item);
        }

        // 3. Capture Total Internal Time
        long totalTime = System.currentTimeMillis() - start;
        System.out.println("Test 3 Total Internal Time: " + totalTime + " ms");

        return new ArrayList<>(userMap.values());
    }
}