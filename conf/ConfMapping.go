package conf

// projecct

var PROJECT_PATH	string
var G_LOG_LEVEL		int

// Server
const SERVER_VERSION 		= "version 1.0"
const SERVER_LISTEN_IP 		= "127.0.0.1"
const SERVER_LISTEN_PORT 	= "7736"
const SERVER_IP_VERSION		= "tcp4"

// Connection

const MAX_CONNECT_LIMIT 	= 1024

// Message

const MAX_MESSAGE_LEN		= 65535

// Task Queue

const TASK_QUEUE_SIZE		= 10
const TASK_QUEUE_BUFFER		= 1024

// Log

const LOG_LEVEL_DEBUG		= 0
const LOG_LEVEL_INFO		= 1
const LOG_LEVEL_WARN		= 2
const LOG_LEVEL_ERROR		= 3
const LOG_LEVEL_CORE_ERROR	= 4