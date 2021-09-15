# Url Shortener - Nginx(Proxy e LoadBalancer) + Golang + PHP
![url shortener@2x (2)](https://user-images.githubusercontent.com/6461792/133510156-ca7f2990-74f1-4e46-b3b1-9ee9342ed73c.png)

## 💻 Projeto
Mini projeto para estudos estilo Bit.ly com Nginx + Proxy. Foi criada uma API em GO para salvar e redirecionar a url. No app PHP faz um requisição onde é gerado as urls aleatoriamente. Como fins didáticos não resolvi criar uma aplicação frontend. Tem um index.php que faz a requisição na API GO. Cada requisição vai cair em algum servidor no Nginx para o loadbalancer.

- PHP para fazer o request na API em GO
- GO para a API
- Nginx servidor e proxy reverso
- K6s para fazer teste de stress

### 🖥️ Iniciando o  projeto
1. ``docker-compose up -d``
2. ``k6 run --vus 1000 --duration 10s /k6s/index.js``


O endpoint principal é: http://localhost:8521/


### 🛠️ Melhorias e pontos para pensar
1. Usar redis para cache para pesquisar se o hash da url já existe no banco
2. Colocar um proxy na api em GO?!
