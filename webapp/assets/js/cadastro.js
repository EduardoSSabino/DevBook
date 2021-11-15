$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(evento){
    evento.preventDefault();
        console.log("Dentro da função usuário");

        if ($('#senha').val() != $('#confirmar-senha').val()){// senha != confirmarSenha
            Swal.fire("Ops...", "As senhas não coincidem!", "error");
            return;
        }

        $.ajax({
            url: "/usuarios",
            method: "POST",
            data: {
                nome: $('#nome').val(),
                email: $('#email').val(),
                nick: $('#nick').val(),
                senha: $('#senha').val()

            }
        }).done(function(retorno){// sucesso
            Swal.fire("Sucesso!", "Usuário cadastrado com sucesso!", "sucess")
            .then(function(){
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }
                }).done(function(){
                    window.location = "/home";
                }).fail(function(){
                    Swal.fire("Ops...", "Erro ao autenticar o usuário!", "error");
                })
            })
        }).fail(function(erro){// falha
            Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
        });
}