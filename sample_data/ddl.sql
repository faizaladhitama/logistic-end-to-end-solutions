-- swift.inbound definition

CREATE TABLE `inbound` (
  `inbound_time` datetime NOT NULL,
  `new_sku_value` varchar(20) CHARACTER SET utf8 NOT NULL,
  `new_title` varchar(63) CHARACTER SET utf8 NOT NULL,
  `expiry_date` varchar(19) CHARACTER SET utf8 NOT NULL,
  `qty` decimal(4,1) NOT NULL,
  `buy_price` decimal(6,1) NOT NULL,
  `total` int(11) NOT NULL,
  `purchase_order_number` varchar(14) CHARACTER SET utf8 NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8mb4;


-- swift.orders definition

CREATE TABLE `orders` (
  `code` varchar(30) NOT NULL,
  `reference_number` varchar(100) NOT NULL,
  `status` varchar(100) NOT NULL,
  `address` text NOT NULL,
  `lat` double NOT NULL,
  `longt` double NOT NULL,
  `timestamp` datetime NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `courier_name` varchar(100) NOT NULL,
  `courier_type` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;


-- swift.outbound definition

CREATE TABLE `outbound` (
  `timestamp` datetime NOT NULL,
  `new_sku` varchar(20) CHARACTER SET utf8 NOT NULL,
  `new_title` varchar(63) CHARACTER SET utf8 NOT NULL,
  `quantity` decimal(2,1) NOT NULL,
  `price` decimal(7,1) NOT NULL,
  `total` int(11) NOT NULL,
  `use_case` varchar(11) CHARACTER SET utf8 NOT NULL,
  `reference_number` varchar(50) CHARACTER SET utf8 DEFAULT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4263 DEFAULT CHARSET=utf8mb4;


-- swift.product definition

CREATE TABLE `product` (
  `sku` varchar(20) CHARACTER SET utf8 NOT NULL,
  `name` varchar(63) CHARACTER SET utf8 NOT NULL,
  `expirable` tinyint(1) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8mb4;


-- swift.stock definition

CREATE TABLE `stock` (
  `sku` varchar(20) CHARACTER SET utf8 NOT NULL,
  `name` varchar(63) CHARACTER SET utf8 NOT NULL,
  `expiry_date` varchar(19) CHARACTER SET utf8 NOT NULL,
  `inbound_date` datetime NOT NULL,
  `inbound_quantity` decimal(4,1) NOT NULL,
  `current_stock` decimal(3,1) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=115 DEFAULT CHARSET=utf8mb4;