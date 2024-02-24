DROP TABLE IF EXISTS TICKET;
DROP TABLE IF EXISTS POSTER;
DROP TABLE IF EXISTS EVENT;
DROP TABLE IF EXISTS PLACE;
DROP TABLE IF EXISTS COMIC;
DROP TABLE IF EXISTS PERSON;
DROP TABLE IF EXISTS IMAGE;

CREATE TABLE Image(
  id SERIAL PRIMARY KEY NOT NULL,
  path TEXT UNIQUE NOT NULL
);

CREATE TABLE Person(
  id SERIAL PRIMARY KEY NOT NULL,
  login TEXT UNIQUE NOT NULL,
  email TEXT UNIQUE NOT NULL,
  salt TEXT NOT NULL DEFAULT 'ya4tr78gf78ert8',
  hash TEXT NOT NULL,
  role INTEGER NOT NULL DEFAULT 3
);

CREATE TABLE Comic(
  id SERIAL PRIMARY KEY NOT NULL ,
  image TEXT DEFAULT '',
  name TEXT NOT NULL ,
  city TEXT DEFAULT 'Столица юмора',
  sentence TEXT DEFAULT '',
  count_kek INTEGER DEFAULT 0 CHECK ( count_kek >= 0)
);

CREATE TABLE Place(
  id SERIAL PRIMARY KEY NOT NULL,
  name TEXT UNIQUE NOT NULL,
  count_ticket INTEGER NOT NULL CHECK ( count_ticket >= 0 )
);

CREATE TABLE Event(
  id SERIAL PRIMARY KEY NOT NULL,
  name TEXT UNIQUE NOT NULL,
  about TEXT DEFAULT ''
);

CREATE TABLE Poster(
  id SERIAL PRIMARY KEY NOT NULL,
  image TEXT DEFAULT '',
  id_event INTEGER REFERENCES Event(id) NOT NULL,
  id_plate INTEGER REFERENCES Place(id) NOT NULL,
  date TIMESTAMP NOT NULL
);

CREATE TABLE Ticket(
  id SERIAL PRIMARY KEY NOT NULL,
  id_poster INTEGER REFERENCES Poster(id) NOT NULL,
  id_person INTEGER REFERENCES Person(id) NOT NULL
);

INSERT INTO event (name, about) VALUES ('Big Stand Up', 'Big Stand Up — шоу с самым большим процентом смеющихся людей. Здесь только опытные комики и шутки, проверенные не одной сотней избирательных зрителей. Приходите убедиться в пятницу, субботу и воскресенье, если вам больше 18 лет и вы свободны в пятницу, субботу и воскресенье.');
INSERT INTO event (name, about) VALUES ('Жёсткий стендап', 'Жёсткий стендап — это шоу, где комики могут шутить обо всем, о чем хотят, не боясь, что их сочтут сумасшедшими. А зрители могут смеяться над всем, над чем хотят, не боясь, что это неуместно. Точно будут шутки про ХХX, не*******ю и к****c. В шоу участвуют 4 комика и ведущий. Состав обновляется раз в месяц — успеете попасть на любимого стендапера. Приходите на шоу, чтобы посмеяться без стыда и, возможно, расширить границы дозволенного. 18+');
INSERT INTO event (name, about) VALUES ('Женщины-комики', 'Женщины-комики — шоу, которое покажет силу женского юмора. В нём участвуют только девушки и только с лучшим своим материалом. Три опытных комикессы, а также ведущая, расскажут качественные шутки. Как про мужчин, феминизм и психотерапию, так и про глобализацию, энтропию и многое другое. Берите подругу, друга, да всех берите и приходите. Мы докажем вам, что женщины умеют шутить обо всём. 18+');
INSERT INTO event (name, about) VALUES ('Стендап Лайнап', 'Лайнап — это стендап-марафон без ведущего. Пять комиков рассказывают по 13 минут качественного материала. В шоу вы можете увидеть опытных стендаперов с разным юмором. Также, в Лайнапе участвуют пока ещё не очень медийные комики, но уже заработавшие себе авторитет в юмористическом комьюнити. Они покажут себя, познакомят с разными формами юмора и, возможно, станут вашими любимчиками. Фишка шоу: если один комик не понравится, через 10 минут уже выступит следующий с другой подачей и отличающимся типом комедии. 18+');
INSERT INTO event (name, about) VALUES ('Сольные концерты', 'Сольные концерты любимых комиков');
INSERT INTO event (name, about) VALUES ('Что у вас случилось?', 'Что у вас случилось? — юмористическо-терапевтическое шоу, в котором комики обсуждают проблемы со зрителями из зала. Три комика вызывают по очереди на сцену человека из зала, который согласен обсудить свой запрос. Комики-терапевты совместно со зрителем-пациентом решают его проблему. 18+');
INSERT INTO event (name, about) VALUES ('Разгоны Офлайн', 'Разгоны Офлайн — это то же самое, что Разгоны Онлайн, только без камер. Это позволяет комикам быть откровеннее в кругу своих друзей и полного зала зрителей. Приходите услышать шутки в их первозданном виде!');
INSERT INTO event (name, about) VALUES ('Вечера комедии', 'Вечера комедии — это шоу, на котором можно услышать современный и актуальный юмор от хороших комиков. Ведь здесь комики показывают, каким должен быть стендап в 2024 году. (Опережают время). Приходите на Вечера комедии, если хотите посмеяться и комфортно провести время, пис. 18+');
INSERT INTO event (name, about) VALUES ('Пятница 13ко', 'Вечернее шоу, в котором собраны все лучшие моменты разных форматов: стендап, лайф интервью и игры с гостями при зрителях. Ведущая шоу Ярослава Тринадцатко и её соведущий Андрей Салеев приглашают в гости звёзд стендап клуба и шоубизнеса, обсуждают последние события в их жизни и смешные моменты из их соцсетей. 18+');
INSERT INTO event (name, about) VALUES ('Secret show', 'Secret Show — это уникальный формат шоу, в котором шоуранеры играют с любопытством и азартом зрителей. Минимум информации, интрига и мистери бокс. Каждое шоу комики готоваят индивидуально, но его суть будет известна только, когда вы зайдёте в зал. Если хотите окунуться в атмосферу закрытых вечеринок или соприкоснуться с чем-то таинственным, то добро пожаловать на Secret Show. 18+');
INSERT INTO event (name, about) VALUES ('На одну тему', 'На одну тему — это офлайн-подкаст с Димой Ковалем и Колей Андреевым, а теперь и с Вовой Бухаровым на определенную и большую тему. Ведущие обсуждают друг с другом, а иногда и с залом, самые разные аспекты в привычном им комедийном разрезе, несмотря на серьёзность заданной темы. Иногда залетают гости. 18+');
INSERT INTO event (name, about) VALUES ('Стендап за 60 секунд', 'Интерактивное стендап-шоу, в котором у каждого участника есть только 1 минута на выступление, после чего он остается на сцене и общается с ведущими. Непредсказуемые интервью, никакой цензуры и много комиков за раз 18+');
INSERT INTO event (name, about) VALUES ('Комедийный квиз', 'Комедийный квиз – это юмористически-познавательная викторина, в которой команды зрителей играют против команды комиков: отвечают на вопросы, соревнуются и выигрывают призы. 18+
Зарегистрируй команду от 2 до 6 человек, чтобы поучаствовать в квизе. Если нет команды, то в графе название команды укажи «я один» или «я одна», и мы найдём тебе команду!
Каждый месяц мы готовим для вас игру. Это значит, что в течение сентября будет одна и та же викторина. Понравилось играть? Приходи в с следующий месяц, там будет новая игра, а, соответственно, новые вопросы!');
INSERT INTO event (name, about) VALUES ('Благотворительный стендап', 'Благотворительный стендап – мероприятие в привычном всем формате стендап-шоу, на котором опытные комики выступают в поддержку самых разных благотворительных фондов. Это шоу, организанованное Соней Медовщиковой, цель которого – привлечь внимание зрителей стендапа к благотворительности и дать возможность заниматься этим с любимыми комиками. Каждое воскресенье новый состав выступающих и новый фонд. Вырученные средства с продажи билетов будут отправляться в виде пожертвования в благотворительные организации.');
INSERT INTO event (name, about) VALUES ('Импров в Стендап Клубе', 'Импров в Стендап Клубе — это театрально-комедийное шоу, которое рождается у вас на глазах. Зритель даёт всего одно слово, а комики на ходу делают из него полноценные комедийные сцены. Куда это может зайти? Не знает никто. Но одно известно точно — двух одинаковых шоу не бывает. Приходите не только посмотреть скетчи от импровизаторов, но и лично принять участие в создании комедии. 18+');
INSERT INTO event (name, about) VALUES ('Альтернативная комедия', 'Альтернативная комедия — это шоу с эксцентричными, по хорошему странными и необычными комиками. На сцене они создают театральную атмосферу, в которой исполняют музыкальные номера, играют персонажей, используют реквизит. Иногда уже известный широкой аудитории комик выступает с необычным для себя форматом и раскрывается для зрителя с новой стороны. На вечере в разных амплуа выступают четыре комика и ведущий. До встречи на этом комедийном кабаре.');
INSERT INTO event (name, about) VALUES ('Энтропия молекулярного котан', 'Энтропия молекулярного котангенса (не очень философско-научный стендап). Стендап, конечно, не очень философский, но зато чуть-чуть научный. Хотя что такое вообще философия? Попробуем нарисовать график… Старое шоу в абсолютно новом формате. Свои вопросы/гипотезы/теории можно присылать на почту emk@standupclub.ru');
INSERT INTO event (name, about) VALUES ('Проверка материала', 'Опытные комики проверяют новые шутки и работают над старыми.18+');
INSERT INTO event (name, about) VALUES ('Ламповый подкаст о философии', 'Ламповый подкаст о философии. Всё как в интернете, но теперь ещё и в коммуникации с залом! Комик Сева Ловкачев и философ Женя Цуркан берут тему с полочки концепций и обсуждают её. Женя отвечает за знания, Сева следит, чтобы всем было смешно и понятно. Строго для зрителей с живым и пытливым умом. 18+');

SELECT * FROM EVENT;
