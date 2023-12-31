// const fetch = require('node-fetch');

const OPENAI_API_KEY = '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';

const headers = {
  'Content-Type': 'application/json',
  'Authorization': `Bearer ${OPENAI_API_KEY}`
};

const data = {
  model: 'gpt-3.5-turbo',
  messages: [
    {
      role: 'system',
      content: 'You are a poetic assistant, skilled in explaining complex programming concepts with creative flair.'
    },
    {
      role: 'user',
      content: '请帮我解释下什么是动态规划？'
    }
  ]
};

fetch('https://api.openai.com/v1/chat/completions', {
  method: 'POST',
  headers: headers,
  body: JSON.stringify(data)
})
  .then(response => response.json())
  .then(result => {
      console.log(result)
      console.log(result.choices[0].message.content)
    // 处理返回的结果
  })
  .catch(error => {
    console.error(error);
    // 处理错误
  });
