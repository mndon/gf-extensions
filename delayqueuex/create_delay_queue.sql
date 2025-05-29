CREATE TABLE delay_queue
(
    id             BIGINT PRIMARY KEY AUTO_INCREMENT,
    topic          VARCHAR(64) NOT NULL COMMENT '队列主题/分类',
    body           TEXT        NOT NULL COMMENT '任务内容，通常为JSON',
    delay_duration INT         NOT NULL COMMENT '延迟时长(秒)',
    ready_time   datetime(6) DEFAULT NULL COMMENT '就绪时间(创建时间 + 延迟时间)',
    status         TINYINT     NOT NULL DEFAULT 0 COMMENT '状态: 0-等待中, 1-处理中, 2-已完成, 3-已取消',
    retry_count    INT         NOT NULL DEFAULT 0 COMMENT '重试次数',
    `updated_time` datetime(6) DEFAULT NULL COMMENT '更新时间',
    `created_time` datetime(6) DEFAULT NULL COMMENT '添加时间',
    INDEX          idx_ready_time (ready_time)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='延迟队列';