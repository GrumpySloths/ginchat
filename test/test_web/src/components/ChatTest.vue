<template>
    <div class="chat-container">
        <div class="chat-messages">
            <div v-for="message in messageHistory" :key="message.id" class="message"
                :class="{ 'sent': message.sender === 'me', 'received': message.sender !== 'me' }">
                <div class="message-content">
                    {{ message.text }}
                </div>
            </div>
        </div>
    </div>
</template>
  
<script>
export default {
    data() {
        return {
            messageHistory: [
                { id: 1, sender: 'me', text: 'Hello' },
                { id: 2, sender: 'friend', text: 'Hi' },
                { id: 3, sender: 'me', text: 'How are you?' },
                { id: 4, sender: 'friend', text: 'I am good, thanks!' },
            ],
        };
    },
};
</script>
  
<style>
.chat-container {
    width: 100%;
    height: 100%;
    display: flex;
    flex-direction: column;
}

.chat-messages {
    flex: 1;
    overflow-y: auto;
}

.message {
    display: flex;
    justify-content: flex-start;
    align-items: center;
    margin-bottom: 10px;
}

.message.sent {
    justify-content: flex-end;
}

.message-content {
    padding: 10px;
    border-radius: 5px;
    background-color: #f0f0f0;
}

.message.sent .message-content {
    background-color: #dcf8c6;
}

.message.received .message-content {
    background-color: #fff;
}
</style>
  