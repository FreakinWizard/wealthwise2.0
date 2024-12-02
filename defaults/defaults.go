package defaults

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Defaults(db *sql.DB) {

	block1Title := `О нас`

	Block1Content1 := `WealthWise — ведущая бухгалтерская компания,
	 работающая на рынке более 3 лет. Мы успешно 
	 обслуживаем более 80 клиентов и имеем более 
	 350 решенных кейсов. Наша компания стремится
	 к максимальной эффективности и качеству 
	 предоставляемых услуг, предлагая
	 индивидуальный подход к каждому клиенту.`

	Block1Content2 := `В нашей команде 12 сотрудников — 
	профессионалы своего дела, обладающие 
	глубокими знаниями и значительным опытом в 
	сфере бухгалтерского учета и финансового анализа.`

	Block1ButtonContent := `Свяжитесь с нами`

	Block1Query := `INSERT INTO Blocks (title, content1, content2, button_content) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(Block1Query, block1Title, Block1Content1, Block1Content2, Block1ButtonContent)
	if err != nil {
		log.Fatal(err)
	}

	block2Title := `Наши преимущества`

	Block2Content1 := `Наша команда состоит из сертифицированных бухгалтеров и аудиторов 
	с многолетним опытом в различных отраслях. 
	Мы предлагаем персонализированные решения, 
	учитывающие потребности каждого клиента, 
	гарантируя безопасность данных и конфиденциальность.`

	Block2Content2 := `Используя передовые программные решения, 
	мы обеспечиваем высокий уровень точности и эффективности, 
	помогая вам сэкономить время и ресурсы для концентрации 
	на основном бизнесе.`

	Block2ButtonContent := `Свяжитесь с нами`

	Block2Query := `INSERT INTO Blocks (title, content1, content2, button_content) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(Block2Query, block2Title, Block2Content1, Block2Content2, Block2ButtonContent)
	if err != nil {
		log.Fatal(err)
	}

	Card1Title := `Бухгалтерский учет Аутсорс`
	Card2Title := `Управленчиский учет`
	Card3Title := `Кадровый учет`
	Card4Title := `Финансовый аудит`
	Card5Title := `Восстановление бухгалтерского учета`
	Card6Title := `Налоговое консультирование и правовые услуги`
	Card1Content := `WealthWise предлагает комплексные бухгалтерские услуги на аутсорсинг, которые помогут вашему бизнесу оптимизировать финансовые процессы и снизить налоговые риски.`
	Card2Content := `WealthWise предлагает услугу управленческого учета, обеспечивая руководство компании актуальной финансовой аналитикой для принятия обоснованных стратегических решений.`
	Card3Content := `WealthWise предоставляет услуги по ведению кадрового учета и точному расчету заработной платы, включая все необходимые налоговые отчисления и взаимодействие с государственными органами.`
	Card4Content := `это комплексная проверка финансовой деятельности компании, целью которой является подтверждение достоверности финансовой отчетности и соответствие операций требованиям законодательства и внутренним стандартам компании.`
	Card5Content := `это услуга по восстановлению корректного учета финансовых операций компании за определенный период, когда учетные данные были утрачены, искажены или велись с ошибками.`
	Card6Content := `Оптимизация налоговых платежей и подготовка отчетности с учетом последних изменений в законодательстве.`

	Card1Query := `INSERT INTO Cards (title, content) VALUES ($1, $2)`
	_, err = db.Exec(Card1Query, Card1Title, Card1Content)
	if err != nil {
		log.Fatal(err)
	}
	Card2Query := `INSERT INTO Cards (title, content) VALUES ($1, $2)`
	_, err = db.Exec(Card2Query, Card2Title, Card2Content)
	if err != nil {
		log.Fatal(err)
	}
	Card3Query := `INSERT INTO Cards (title, content) VALUES ($1, $2)`
	_, err = db.Exec(Card3Query, Card3Title, Card3Content)
	if err != nil {
		log.Fatal(err)
	}
	Card4Query := `INSERT INTO Cards (title, content) VALUES ($1, $2)`
	_, err = db.Exec(Card4Query, Card4Title, Card4Content)
	if err != nil {
		log.Fatal(err)
	}
	Card5Query := `INSERT INTO Cards (title, content) VALUES ($1, $2)`
	_, err = db.Exec(Card5Query, Card5Title, Card5Content)
	if err != nil {
		log.Fatal(err)
	}
	Card6Query := `INSERT INTO Cards (title, content) VALUES ($1, $2)`
	_, err = db.Exec(Card6Query, Card6Title, Card6Content)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateTable(db *sql.DB) {
	// SQL query to create the Blocks table if it doesn't already exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS Blocks (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		content1 TEXT NOT NULL,
		content2 TEXT NOT NULL,
		button_content TEXT NOT NULL
	)`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}

func CreateCardTable(db *sql.DB) {
	// SQL query to create the Blocks table if it doesn't already exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS Cards (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL,
		content TEXT NOT NULL
	)`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Failed to create table: ", err)
	}
}

func CreateImageTable(db *sql.DB) {
	query :=
		`CREATE TABLE IF NOT EXISTS images (
		id SERIAL PRIMARY KEY,
		image_url TEXT
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal("Failed to create image table: ", err)
	}
}
