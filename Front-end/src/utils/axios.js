let token = ""
if (localStorage.getItem("token") != null) {
    token = localStorage.getItem("token").split('"')[1];
}
else{
    token = ""
}

const axios = require("axios");
const axiosInstance = axios.create({
    baseURL: 'http://localhost:8080/',
    headers: {
        Authorization: `bearer ${token}`,
      },
});
axiosInstance.interceptors.request.use(
  (config) => new Promise((resolve) => setTimeout(() => resolve(config), 700))
);

export default axiosInstance;