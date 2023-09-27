CREATE TABLE if not exists Wildberry (
  "order_uid" text ,
  "track_number" text,
  "entry" text,
  "delivery" text ,
  "payment" text ,
  "items" json,
  "locale" text,
  "internal_signature" text,
  "customer_id" text,
  "delivery_service" text,
  "shardkey" text,
  "sm_id" bigint,
  "date_created" text,
  "oof_shard" text,
	FOREIGN KEY (delivery) REFERENCES Deliveryt (phone) ON DELETE CASCADE,
	FOREIGN KEY (payment) REFERENCES Paymentt (transaction) ON DELETE CASCADE
)

Create table if not exists Deliveryt(
  "name" text ,
  "phone" text Primary key,
  "zip" text,
  "city" text,
  "address" text,
  "region" text,
  "email" text

  )
  
  Create table if not exists Paymentt(
  "transaction" text Primary key,
  "request_id" text,
  "currency" text,
  "provider" text,
  "amount" bigint,
  "payment_dt" bigint,
  "bank" text,
  "delivery_cost" bigint,
  "goods_total" bigint,
  "custom_fee" bigint
	 
  )