package apiserver

import "net/http"

//responseWriter Наследуется от ResponseWriter, а значит удовлетворяет тем же интерфейсам.
//отдельный код нужен для того чтобы мидлвейр logRequest мог получить код результата обработки следующих хендлеров
//и вывести его в логи
type responseWriter struct {
	http.ResponseWriter
	code int
}

//WriteHeader переопределяем метод, чтобы он сначала записывал код ответа в наше кастомное поле, а потом уже вызывал
//стандартный метод экспортируемого интерфейса.
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.code = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
