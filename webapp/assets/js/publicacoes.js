$('#nova-publicacao').on('submit', criarPublicacao);
$(document).on('click', '.curtir-publicacao', curtirPublicacao);
$(document).on('click', '.descurtir-publicacao', descurtirPublicacao);
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$(' .deletar-publicao').on('click', deletarPublicacao);
function criarPublicacao(evento){
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function(){
        window.location = "/home";
    }).fail(function(){
        Swal.fire("Ops...", "Erro ao criar publicação!", "error");
    })
}

function curtirPublicacao(evento){
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closet('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);// desabilitando o botao de curtir pra evitar a sobrecarga de requisições caso alguem clique disparadamente

    $.ajax({
        url: `/publicacoes/${publicacaoId}/curtir`,
        method: "POST"
    }).done(function(){
       const contadorDeCurtidas = elementoClicado.Next('span');
       const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

       contadorDeCurtidas.text(quantidadeDeCurtidas + 1);

       elementoClicado.addClass('descurtir-ooublicacao');
       elementoClicado.addClass('text-denger'); // deixando o coração vermelho depois de curtir
       elementoClicado.removeClass('curtir-publicacao');

    }).fail(function(){
        Swal.fire("Ops...", "Erro ao curtir a publicação!", "error");
    }).always(function(){
        elementoClicado.prop('disabled', false);// habilitando o botão novamente
    })
}

function descurtirPublicacao(evento){
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closet('div').data('publicacao-id');

    elementoClicado.prop('disabled', true);// desabilitando o botao de curtir pra evitar a sobrecarga de requisições caso alguem clique disparadamente

    $.ajax({
        url: `/publicacoes/${publicacaoId}/descurtir`,
        method: "POST"
    }).done(function(){
       const contadorDeCurtidas = elementoClicado.Next('span');
       const quantidadeDeCurtidas = parseInt(contadorDeCurtidas.text());

       contadorDeCurtidas.text(quantidadeDeCurtidas - 1);

       elementoClicado.removeClass('descurtir-ooublicacao');
       elementoClicado.removeClass('text-denger'); // deixando o coração vermelho depois de curtir
       elementoClicado.addClass('curtir-publicacao');// habilita "curtir" como minha próxima função

    }).fail(function(){
        Swal.fire("Ops...", "Erro ao curtir a publicação", "error");
    }).always(function(){
        elementoClicado.prop('disabled', false);// habilitando o botão novamente
    })
}

function atualizarPublicacao(){
$(this).prop('disabled', false);// desabilitando um botão que foi clicado

const publicacaoId = $(this).data('publicacao-id');
$.ajax({
    url: `/publicacoes/${publicacaoId}`,
    method: "PUT",
    data: {
        titulo: $('#titulo').val(),
        conteudo: $('#conteudo').val()
    }
}).done(function(){
    Swal.fire(
        'Sucesso!', 'Publicação criada com sucesso!', 'sucess')
        .then(function(){
            window.location = "/home";
        })
}).fail(function(){
    Swal.fire("Ops...", "Erro ao editar a publicação!", "error");
}).always(function(){
    $('#atualizar-publicacao').prop('disabled', false);// habilitando o botão
    })
}

function deletarPublicacao(evento){
    evento.preventDefault();

    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluitr essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "cancelar",
        icon: "warning"
    }).then(function(confirmacao){
            if (!confirmacao.value) return;
   

    const elementoClicado = $(evento.target);
    const publicacao = elementoClicado.closest('div');
    const publicacaoId = publicacao.data('publicacao-id');

    elementoClicado.prop('disabled', true);// desabilitando o botao de curtir pra evitar a sobrecarga de requisições caso alguem clique disparadamente
    
    $.ajax({
        url: `/publicacoes/${publicacaoId}`,
        method: "DELETE"
    }).done(function(){
        publicacao.fadeOut("slow", function(){
            $(this).remove();
        });
    }).fil(function(){
        Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
    });
 })
}