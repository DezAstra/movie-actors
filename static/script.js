$(document).ready(function() {
    // загрузка актёров на экран
    fetchActors();

    // Загрузка списка национальности
    fetchNationalities();

    // Обработчик поиска
    $('#search').on('input', function() {
        fetchActors();
    });

    // Обработчик фильтра по гонорарам
    $('#earnings-min, #earnings-max').on('input', function() {
        fetchActors();
    });

    // Обработчик фильтра по количеству фильмов
    $('#films-min, #films-max').on('input', function() {
        fetchActors();
    });

    // Обработчик фильтра по национальности
    $('#nationality').on('change', function() {
        fetchActors();
    });

    // Функция  получения списка актёров с фильтрами
    function fetchActors() {
        let search = $('#search').val();
        let filmsMin = $('#films-min').val();
        let filmsMax = $('#films-max').val();
        let earningsMin = $('#earnings-min').val();
        let earningsMax = $('#earnings-max').val();
        let nationality = $('#nationality').val();


        let url = `/actors?search=${search}`;

        if (earningsMin) {
            url += `&earnings_min=${earningsMin}`;
        }
        if (filmsMin) {
            url += `&films_min=${filmsMin}`;
        }
        if (filmsMax) {
            url += `&films_max=${filmsMax}`;
        }
        if (earningsMax) {
            url += `&earnings_max=${earningsMax}`;
        }
        if (nationality) {
            url += `&nationality=${nationality}`;
        }

        console.log("URL:", url);

        $.get(url, function(data) {
            $('#actor-list').empty();
            data.forEach(actor => {
                $('#actor-list').append(`
                    <div class="actor-item">
                        <h3>${actor.first_name} ${actor.last_name}</h3>
                        <p><strong>Национальность:</strong> ${actor.nationality || 'Не указана'}</p>
                        <p><strong>Фильмов:</strong> ${actor.film_count}</p>
                        <p><strong>Гонорары:</strong> $${actor.total_earnings} млн</p>
                    </div>
                `);
            });
        });
    }

    // Функция получения списка национальностей
    function fetchNationalities() {
        $.get("/nationalities", function(data) {
            let nationalitySelect = $('#nationality');
            nationalitySelect.empty();
            nationalitySelect.append('<option value="">Выберите национальность</option>');
            data.forEach(nat => {
                nationalitySelect.append(`<option value="${nat.id}">${nat.name}</option>`);
            });
        });
    }
});
