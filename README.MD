# Entities, Attribute, Relationship
* ## Entities & Attribute
    * ## Product
        * Product ID
        * SKU
        * Name
        * Is Expireable
        * Created Date
        * Updated Date
    * ## Stock
        * Stock ID
        * SKU
        * Name
        * Expiry Date
        * Inbound Date
        * Inbound Quantity
        * Current Stock
        * Created Date
        * Updated Date
    * ## Inbound
        * Inbound ID
        * Inbound Time
        * Product Name
        * Product SKU
        * Expiry Date
        * Qty
        * Buy Price
        * Total
        * Purchase Order Number
        * Created Date
        * Updated Date
    * ## Outbound
        * Outbound ID 
        * Reference Number
        * Timestamp
        * Product SKU
        * Product Name
        * Qty
        * Price
        * Total
        * Use Case
        * Created Date
        * Updated Date
    * ## Order
        * Order ID
        * Client Reference Number Order ID
        * Status
        * Created Date
        * Updated Date
    * ## Order History
        * Order History ID   
        * Order ID
        * Status
        * Created Date
        * Updated Date
    * ## Logistic
        * Logistic ID
        * Order ID
        * Courier Name
        * Courier Type
        * Distance
        * Price
        * Created Date
        * Updated Date
    * ## Product Attribute
        * SKU
        * Weight
        * Height
        * Width
        * Length
        * Is Fresh Product
        * Created Date
        * Updated Date
* ## Relationship
    * Product <-> Stock <-> Inbound <-> Outbound <-> Product Attribute : SKU field
    * Stock <-> Inbound : Inbound date & qty field
    * Order <-> Order History <-> Logistic : Order ID field
    
# Subsystems / Services
* Order System
    * Create, Read, Update Order
    * Connect with third party logistic
    * Publish stock reduction to Inventory System
* Inventory System
    * Product Information
    * Stock Information
    * Inbound & Outbound Information

# Tech Stack
* [React] - Web framework to create robust and reusable component
* [Echo] - Web framework to create fast API
* [Kafka] - Can be used as message broker. Send notification about order and real-time stock update
* [Redis] - In-memory database for storing cache
* [MySql] - Database
* [Varnish] - HTTP Accelerator
* [Cloudflare CDN] - CDN
* [Jenkins] - CI/CD
* [Gitlab] - Git Repository
* [Nginx] - Custom load balancer (multi node deployment only)
* [AWS] - Cloud Service Provider (Kafka, Compute Engine, Docker, Kubernetes, etc)

[React]: <https://reactjs.org/>
[Echo]: <https://echo.labstack.com/>
[Kafka]: <https://kafka.apache.org/>
[Redis]: <https://redis.io/>
[MySql]: <https://www.mysql.com/>
[Varnish]: <https://varnish-cache.org/>
[Cloudflare CDN]: <https://www.cloudflare.com/cdn/>
[Jenkins]: <https://www.jenkins.io/>
[Gitlab]: <https://about.gitlab.com/>
[Nginx]: <https://www.nginx.com/>
[AWS]: <https://aws.amazon.com/>