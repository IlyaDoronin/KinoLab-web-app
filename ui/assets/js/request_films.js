function $(el){
    return document.querySelector(el)
}


// Колличество страниц с фильмами 
// let pageCount
// fetch(`http://${HOST}/pageCount?table=film`, {
//         "Access-Control-Allow-Methods": "GET,PUT,POST,DELETE,PATCH,OPTIONS",
//         "Access-Control-Allow-Origin": "*"
//     }).then( response =>  response.json() )
//     .then(count => pageCount = count.page_count )


// текущая страница для загрузки комментариев(комменты в api по страницам)
let filmPage = 1

btn = $('#more-btn')
previewList = $('.wrapper__preview')

// Ajax запрос новых фильмов
function getFilms(){
    // Подсчёт загруженных фильмов для правильного ajax запроса
    number = document.querySelectorAll('.preview').length
    URL = `http://${HOST}/website/films?page=${filmPage}`

    fetch(URL, { 
        method: 'GET',
        headers: { 'Content-Type': 'application/json' }
    }).then(response => {
            if (response.ok){    

                // if( pageCount == filmPage){
                //     btn.style.display = 'none'
                // }

                filmPage++
                return response.json()
            }
            else
                console.error('Ошибка получения данных');
        }).then( films => {            
            films.films.map( film => generatePreview(film))
        })

}

// генерация HTML кода(добавление новых фильмов) 
function generatePreview(film){
    let preview = document.createElement('div')
    preview.classList.add('preview')
    preview.innerHTML = `
    <a href="/film/${film.ID}" class="card__wrapper">
    <div class="preview__info">                        
        <img src="${film.PosterURL}" class="card__pic">
    </div>
    <div class="preview__popup">
        <h4 class="preview__popup-title">${film.FilmName}</h4>
        <h5 class="preview__popup-genre">${film.Genres}</h5>
    </div>
</a>
    `
    previewList.append(preview)
}

btn.addEventListener('click', () => {
    getFilms(filmPage)
})

getFilms(filmPage)
