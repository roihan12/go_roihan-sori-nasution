package category

type categoryUsecase struct {
	categoryRepository Repository
}

func NewCategoryUsecase(cr Repository) Usecase {
	return &categoryUsecase{
		categoryRepository: cr,
	}
}

func (cu *categoryUsecase) GetAll() []Domain {
	return cu.categoryRepository.GetAll()
}

func (cu *categoryUsecase) GetByID(id string) Domain {
	return cu.categoryRepository.GetByID(id)
}

func (cu *categoryUsecase) Create(categoryDomain *Domain) Domain {
	return cu.categoryRepository.Create(categoryDomain)
}

func (cu *categoryUsecase) Update(id string, categoryDomain *Domain) Domain {
	return cu.categoryRepository.Update(id, categoryDomain)
}

func (cu *categoryUsecase) Delete(id string) bool {
	return cu.categoryRepository.Delete(id)
}
