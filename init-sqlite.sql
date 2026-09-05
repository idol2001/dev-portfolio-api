-- 创建默认管理员账号 (SQLite 语法)
-- 用户名: admin
-- 密码: admin123 (bcrypt 加密)

-- 注意: SQLite 使用 INSERT OR IGNORE 替代 MySQL 的 INSERT IGNORE
INSERT OR IGNORE INTO users (username, password, email, role, nickname, created_at, updated_at)
VALUES ('admin', '$2b$14$kuY.g8TZNY117IpvdJ5y5.LDFf0yF7.vnp6wi9BlBh4WmXKbtIS.W', 'admin@example.com', 'admin', '管理员', datetime('now'), datetime('now'));
