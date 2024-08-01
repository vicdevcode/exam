package app

import (
	"context"
	"fmt"

	"github.com/vicdevcode/exam/internal/entity"
	"github.com/vicdevcode/exam/internal/sqlite"
	"github.com/vicdevcode/exam/internal/usecase"
)

func Migrate(runType string, db *sqlite.Sqlite) {
	switch runType {
	case "create":
		if err := create(db); err != nil {
			panic(err)
		}
	case "drop":
		if err := drop(db); err != nil {
			panic(err)
		}
	case "reset":
		if err := drop(db); err != nil {
			panic(err)
		}
		if err := create(db); err != nil {
			panic(err)
		}
	default:
		panic("?")
	}
}

func create(db *sqlite.Sqlite) error {
	if err := db.AutoMigrate(
		&entity.Category{},
		&entity.SubCategory{},
		&entity.Card{},
	); err != nil {
		return err
	}
	return nil
}

func drop(db *sqlite.Sqlite) error {
	tables := []string{"cards", "sub_categories", "categories"}
	for _, t := range tables {
		if err := db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", t)).Error; err != nil {
			return err
		}
	}
	return nil
}

func addData(uc usecase.UseCases) error {
	geo, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Алгебра и геометрия"},
	)
	if err != nil {
		return err
	}
	c2, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Математический анализ"},
	)
	if err != nil {
		return err
	}
	c3, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Дискретная математика"},
	)
	if err != nil {
		return err
	}
	c4, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Вычислительная математика"},
	)
	if err != nil {
		return err
	}
	c5, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Алгоритмы и анализ сложности"},
	)
	if err != nil {
		return err
	}
	c6, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Языки программирования"},
	)
	if err != nil {
		return err
	}
	c7, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Объектно-ориентированное программирование"},
	)
	if err != nil {
		return err
	}
	c8, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Компьютерные сети"},
	)
	if err != nil {
		return err
	}
	c9, err := uc.CategoryUseCase.Create(context.Background(), entity.Category{Name: "Базы данных"})
	if err != nil {
		return err
	}
	c10, err := uc.CategoryUseCase.Create(
		context.Background(),
		entity.Category{Name: "Программная инженерия"},
	)
	if err != nil {
		return err
	}

	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: geo.ID, Name: "Векторы"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: geo.ID, Name: "Матрицы и определители"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: geo.ID, Name: "Прямая и плоскость"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: geo.ID, Name: "Кривые второго порядка"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c2.ID, Name: "Введение"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c2.ID, Name: "Производная и дифференциал"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c2.ID, Name: "Интеграл"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c2.ID, Name: "Ряды"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c2.ID, Name: "Кратные интегралы"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c3.ID, Name: "Комбинаторика"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c3.ID, Name: "Булевы функции"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c3.ID, Name: "Элементы теории графов"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c4.ID, Name: "Численные методы линейной алгребы"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c4.ID, Name: "Численное интегрирование"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c4.ID, Name: "Методы решения нелинейных уравнений"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c4.ID, Name: "Элементы теории разностных схем"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{
			CategoryID: c4.ID,
			Name:       "Численные методы решения задачи Коши для обыкновенных дифференциальных уравнений",
		},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{
			CategoryID: c4.ID,
			Name:       "Разностные методы решения задач математической физики",
		},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c5.ID, Name: "Основы анализа алгоритмов"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c5.ID, Name: "Стратегии алгоритмов"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c5.ID, Name: "Основные алгоритмы"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c6.ID, Name: "Языки программирования"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c6.ID, Name: "Способы описания языков программирования"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c6.ID, Name: "Общее представление о процессе трансляции"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{
			CategoryID: c6.ID,
			Name:       "Промежуточные (внутренние) представления программы",
		},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c6.ID, Name: "Связные структуры данных"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c7.ID, Name: "Основные принципы и конструкции ООП"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c7.ID, Name: ".NET Framework"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c7.ID, Name: "Объектное-ориентированная декомпозиция"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c8.ID, Name: "Локальные сети и сетевое оборудование"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c8.ID, Name: "Эталонная семиуровневая модель ISO/OSI"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c8.ID, Name: "Стандарты Ethernet"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c8.ID, Name: "Сетевой уровень модели OSI"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c8.ID, Name: "Служба DNS"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c8.ID, Name: "Протокол HTTP"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c9.ID, Name: "Модели данных"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{
			CategoryID: c9.ID,
			Name:       "Проектирование реляционных БД на основе принципов нормализации",
		},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c9.ID, Name: "Язык SQL"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c9.ID, Name: "Целостность и моделирование данных"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c10.ID, Name: "Этапы разработки ПО"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c10.ID, Name: "Проектирование ПО"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c10.ID, Name: "Конструирование ПО"},
	)
	if err != nil {
		return err
	}
	_, err = uc.SubCategoryUseCase.Create(
		context.Background(),
		entity.SubCategory{CategoryID: c10.ID, Name: "Тестирование ПО"},
	)
	if err != nil {
		return err
	}
	return nil
}
