
CREATE TABLE mind_product (
    product_id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name text NOT NULL,
    description text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mind_partner (
    partner_id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name text NOT NULL,
    city text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mind_invoice (
    invoice_id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    partner_id int NOT NULL REFERENCES mind_partner(partner_id),
    payment_method_id text NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE mind_invoice_item (
    invoice_item_id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    item_id text NOT NULL REFERENCES mind_voucher(voucher_code),
    item_amount int NOT NULL,
    item_price int NOT NULL,
    item_discount int NOT NULL
);

CREATE TABLE mind_invoice_item_relationship (
    invoice_id int NOT NULL REFERENCES mind_invoice(invoice_id),
    invoice_item_id int NOT NULL REFERENCES mind_invoice_item(invoice_item_id)
);

CREATE TABLE mind_voucher (
    voucher_code text PRIMARY KEY,
    product_id int NOT NULL REFERENCES mind_product(product_id),
    nominal numeric NOT NULL,
    duration_month int NOT NULL,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    expired_at timestamp with time zone,
    activated_at time without time zone
);

CREATE TABLE mind_partner_voucher (
    invoice_id int NOT NULL REFERENCES mind_product(product_id),
    partner_id int NOT NULL REFERENCES mind_partner(partner_id),
    voucher_code text NOT NULL REFERENCES mind_voucher(voucher_code),
    purchase_value numeric NOT NULL
);

Create or replace function random_string(length integer) returns text as
$$
declare
  chars text[] := '{0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z}';
  result text := '';
  i integer := 0;
begin
  if length < 0 then
    raise exception 'Given length cannot be less than 0';
  end if;
  for i in 1..length loop
    result := result || chars[1+random()*(array_length(chars, 1)-1)];
  end loop;
  return result;
end;
$$ language plpgsql;


INSERT INTO mind_voucher (product_id, code, nominal, duration_month, expired_at)
VALUES (1, random_string(8), 100000 , 6, now() + '1 month'::interval * 120)
RETURNING code;

INSERT INTO mind_partner (name, city)
VALUES
 ('Hendro', 'Ujung Batu');

 INSERT INTO mind_product (name, description)
VALUES
 ('Mindtrex Academy', 'for Brunei');

