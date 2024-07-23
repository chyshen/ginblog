## 1. 简介

[Axios](https://www.axios-http.cn/) 是一个基于 [promise](https://javascript.info/promise-basics) 网络请求库，作用于[`node.js`](https://nodejs.org/) 和浏览器中。 它是 [isomorphic](https://www.lullabot.com/articles/what-is-an-isomorphic-application) 的(即同一套代码可以运行在浏览器和`node.js`中)。在服务端它使用原生 `node.js` `http` 模块, 而在客户端 (浏览端) 则使用 `XMLHttpRequests`。



## 2. 安装

```shell
pnpm install axios
```



## 3. Axios实例

> 以`Vue`中为例

```typescript
// 创建src/common/request.ts文件
import axios from 'axios'

export default const instance = axios.create({
  baseURL: 'https://localhost:8080/api/v1/',
  timeout: 1000
});
```



## 4. 拦截器



### 4.1 请求拦截器

```typescript
// src/common/request.ts

instance.interceptors.request.use(function (config) {
    // 在发送请求之前做些什么
    return config;
  }, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
  });
```



### 4.2 响应拦截器

```typescript
// src/common/request.ts

instance.interceptors.response.use(function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    return response;
  }, function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
  });
```



## 5. 实例

```typescript
import axios from 'axios'
import { ElMessage } from 'element-plus'
// 从状态管理（如：Pinia）中导入token
import { useTokenStore } from '@/stores/token'
// 导入Vue路由
import router from '@/router'

export default const instance = axios.create({
  baseURL: 'https://localhost:8080/api/v1/',
  timeout: 1000
});

// 请求拦截器
instance.interceptors.request.use((config) => {
    // 在发送请求之前做些什么
    // 获取token
    const tokenStore = useTokenStore()
    // 判断请求是否添加token
     if (tokenStore.token) {
      config.headers.Authorization = tokenStore.token
    }
    return config
  }, (error) => {
    // 对请求错误做些什么
    return Promise.reject(error);
  });

// 响应拦截器
instance.interceptors.response.use((response) => {
    if (response.data.code === 200) {
        return response.data
    } else {
        // 状态码不是200
        ElMessage.error('当前请求失败：' + response.data.messages ? response.data.message: '请求失败')
        // 异步状态转化成失败的状态
        return Promise.reject(response.data)
    }
  }, (error) => {
    // 未登录响应状态码401，未认证状态401，给出对应的提示，并跳转到登录页（?检查err.response是否存在，防止空值错误）
    if (err.response?.status === 401) {
        ElMessage.error('登录已失效，请重新登录')
      	// 获取token状态
      	const tokenStore = useTokenStore()
      	// 清除token
      	tokenStore.removeToken()
      	// 跳转到登录页
      	router.push('/login')
    }
    if (err.response?.status === 500) {
      	// ElMessage.error('当前无权限访问该接口')
    } else {
      	// 判断err.response是否存在，以避免访问undefined的属性
      	ElMessage.error(err.message || '请求失败')
    }
    	// 异步的状态转化成失败的状态
    	return Promise.reject(err)
  	}
  )
```

