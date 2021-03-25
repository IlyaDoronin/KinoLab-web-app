// Начало выбора параметров( choice )
// Создание параметра
let choices = document.querySelectorAll(".main__filter_item"); // Контейнер с чёйсами
const tagsList = document.querySelector(".filter__tags"); // Список пояляющихся фильтров
const search = document.querySelector('.main__search_input')

for (let i = 0; i < choices.length; i++) {
    createElem(
        choices[i],
        choices[i]
            .closest(".main__filter_choice")
            .getAttribute("data-type"),
    );
}
function createElem(item, type) {
    let text = item.querySelector("p").innerHTML;
    item.querySelector("input").addEventListener("change", function () {
        if (this.checked) {
            const id = this.getAttribute("id");
            let li = document.createElement("li");
            li.className = "filter__tags-item";
            li.setAttribute("data-number", id);
            li.setAttribute("data-type", type);
            li.innerHTML = `<p>${text}</p><span class="filter__tags-item__cross"></span>`;
            li.addEventListener("click", () => {
                document.querySelector(`#${id}`).checked = false;
                document
                    .querySelectorAll(".filter__tags-item")
                    .forEach((item) => {
                        if (item.getAttribute("data-number") == id)
                            item.remove();
                    });
            });
            tagsList.append(li);
        } else {
            removeElem(this);
        }
    });
}
function removeElem(context) {
    const id = context.getAttribute("id");
    document.querySelectorAll(".filter__tags-item").forEach((item) => {
        if (item.getAttribute("data-number") == id) {
            item.remove();
        }
    });
}
// Конец выбора параметров( choice )

// Ajax запрос по фильтру 
function sendFilters(filters){
    fetch(`http://${HOST}/website/filter`, { 
        method: 'POST',
        body: filters
    }).then(response => {
            return response.json()
    }).then( filters => {  
        previewList.innerHTML = ''
        $('#more-btn').style.display = 'none'
        filters.map(filter => generatePreview(filter))
    })
}

// Поиск по параметрам
let filters = {
    search: '',
    genres: [],
    authors: [],
    actors: [],
    years: [],
};
function takeFiltes() {
    filters = {
        search: '',
        genres: [],
        authors: [],
        actors: [],
        years: [],
    };
    // Очитка объедка от старых фильтров
    filters.search = search.value
    tagsList.querySelectorAll(".filter__tags-item").forEach((item) => {
        switch (item.getAttribute("data-type")) {
            case "genre":
                filters.genres.push(item.querySelector("p").innerHTML);
                break;
            case "authors":
                filters.authors.push(item.querySelector("p").innerHTML);
                break;
            case "actor":
                filters.actors.push(item.querySelector("p").innerHTML);
                break;
            case "year":
                filters.years.push(item.querySelector("p").innerHTML);
                break;
        }
    });

    sendFilters(JSON.stringify(filters));
}

document.querySelector(".filter__btn").addEventListener("click", (e) => {
    e.preventDefault();
    takeFiltes();
});

