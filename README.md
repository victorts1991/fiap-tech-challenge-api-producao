# fiap-tech-challenge-api

SONAR_TOKEN

### Tech Challenge 4:
### Passos para homologação dos professores da Fiap

1. Faça o login na plataforma da AWS;
2. Crie um repositório privado no ECR da AWS chamado fiap-tech-challenge-api-producao;
3. Acesse IAM->Usuários e crie um novo usuário chamado Github;
4. Com esse usuário criado, vá até a listagem de usuários e acesse os detalhes do mesmo;
5. No menu Permissões que irá aparecer na tela de detalhes, clique no botão "Adicionar permissões" que aparece no canto direito;
6. Na tela que irá aparecer, selecione a opção "Anexar políticas diretamente";
7. Pesquise pela permissão "AmazonEC2ContainerRegistryPowerUser" e adicione ela;
8. Após isso, de volta a tela de detalhes do usuário, clique na aba "Credenciais de Segurança", e no bloco "Chaves de acesso", clique em "Criar chave de acesso";
9. Na tela que irá se abrir, selecione a opção "Command Line Interface (CLI)" e clique em próximo;
10. No valor da etiqueta, coloque o valor "github actions" ou qualquer um que prefira para identificar posteriormente;
11. Copie os valores dos campos "Chave de acesso" e "Chave de acesso secreta";
12. Na plataforma do Github, acesse o menu "Settings" do projeto, na tela que se abrir, clique no menu Security->Secrets and variables->Actions;
13. Adicione uma "repository secret" chamada AWS_ACCESS_KEY_ID com o valor copiado de "Chave de acesso", e crie outra "repository secret" chamada AWS_SECRET_ACCESS_KEY com o valor copiado de "Chave de acesso secreta";
14. Após isso qualquer commit neste repositório que for para a branch "main", irá subir uma imagem desta api no ECR da AWS;
