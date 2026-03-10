# Delivery Tracker API

We're tracking FLAVOR all the way to your door! Real-time delivery tracking that keeps customers in the loop.

The Delivery Tracker API handles everything that happens after food leaves the kitchen. Driver assignment, GPS tracking, ETA calculations, and delivery confirmation. Customers can watch their food's journey on a live map — because anticipation is half the flavor.

## Features

- **Real-time GPS Tracking** — Driver location updates every few seconds
- **ETA Calculation** — Dynamic ETA based on current location, traffic, and route
- **Delivery Lifecycle** — Track from assigned through picked_up, in_transit, to delivered
- **Driver Management** — Assign and reassign drivers based on proximity and availability

## Why Go?

Like the Grill Master Service, we chose Go for its performance characteristics. Real-time location tracking means processing thousands of GPS updates per second. Go's efficient concurrency model handles this with ease.
