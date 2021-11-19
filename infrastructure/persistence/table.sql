
-- ----------------------------
-- Table structure for stock
-- ----------------------------
DROP TABLE IF EXISTS "public"."stock";
CREATE TABLE "public"."stock" (
  "symbol" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "name" varchar COLLATE "pg_catalog"."default" NOT NULL,
  "net_profit_cagr" float8,
  "north_net_inflow" varchar COLLATE "pg_catalog"."default",
  "ps" float8,
  "type" int8,
  "percent" float8,
  "has_follow" bool,
  "tick_size" float8,
  "pb_ttm" float8,
  "float_shares" int8,
  "current" float8,
  "amplitude" float8,
  "pcf" float8,
  "current_year_percent" float8,
  "float_market_capital" int8,
  "north_net_inflow_time" varchar COLLATE "pg_catalog"."default",
  "market_capital" int8,
  "dividend_yield" int8,
  "lot_size" int8,
  "roe_ttm" float8,
  "total_percent" float8,
  "percent5m" int8,
  "income_cagr" float8,
  "amount" int8,
  "chg" float8,
  "issue_date_ts" int8,
  "eps" float8,
  "main_net_inflows" int8,
  "volume" int8,
  "volume_ratio" float8,
  "pb" float8,
  "followers" int8,
  "turnover_rate" float8,
  "first_percent" float8,
  "pe_ttm" float8,
  "total_shares" int8,
  "limitup_days" int8
)
;

-- ----------------------------
-- Primary Key structure for table stock
-- ----------------------------
ALTER TABLE "public"."stock" ADD CONSTRAINT "stock_pkey" PRIMARY KEY ("symbol");