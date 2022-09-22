package main

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

func (j *JobSite) sendAll() {
	for _, v := range j.subscribers {
		v.handleEvent(j.vacancies)
	}
}

func (j *JobSite) addVacancy(vacancy string) {
	//fmt.Println("Adding vacancy ", vacancy)
	j.vacancies = append(j.vacancies, vacancy)
	//fmt.Println("Vacancies: ", len(j.vacancies))
	j.sendAll()
}

func (j *JobSite) removeVacancy(vacancy string) {
	for i, v := range j.vacancies {
		//fmt.Println("Removing vacancy ", vacancy, " ==== ", v)
		if v == vacancy {
			j.vacancies = append(j.vacancies[:i], j.vacancies[i+1:]...)
		}
	}
	j.sendAll()
}
func (j *JobSite) subscribe(observer Observer) {
	j.subscribers = append(j.subscribers, observer)
}

func (j *JobSite) unsubscribe(observer Observer) {
	for i, v := range j.subscribers {
		if v == observer {
			j.subscribers = append(j.subscribers[:i], j.subscribers[i+1:]...)
		}
	}
}
