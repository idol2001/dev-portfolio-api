-- 创建默认管理员账号
-- 用户名: admin
-- 密码: admin123 (bcrypt 加密)

-- 方式 1: 如果 users 表为空，直接插入
INSERT IGNORE INTO users (username, password, email, role, nickname, created_at, updated_at)
SELECT 'admin', '$2a$14$YJxgZ7z8K9vN2qR5tU3wL.xC1vB8nM4jK6lP0oS2dF7gH9iA3cE5m', 'admin@example.com', 'admin', '管理员', NOW(), NOW()
WHERE NOT EXISTS (SELECT 1 FROM users WHERE username = 'admin');
