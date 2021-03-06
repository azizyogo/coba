package service

import (
	"middleware-mmksi/dsf/metadata/repo"
	"middleware-mmksi/dsf/metadata/response"
	"middleware-mmksi/dsf/metadata/service/request"
)

type DsfProgramService interface {
	GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error)
	GetPackageNames(params request.HeaderPackageNameRequest, queryParams request.ParamsPackageNameRequest) (*response.PackageNameResponse, error)
	GetCarConditions() (*response.CarConditionResponse, error)
	GetPackages(paramHeader request.HeaderPackageRequest, reqBody request.PackageRequest) (*response.PackageResponse, error)
	GetVariants(paramHeader request.HeaderUnitByModelsRequest) (*response.VariantsResponse, error)
	GetPaymentTypes() (*response.PaymentTypesResponse, error)
	GetModels(params request.ModelsRequest) (*response.ModelsResponse, error)
	GetBrands(params request.BrandsRequest) (*response.BrandsResponse, error)
	GetVehicleCategory() (*response.VehicleCategory, error)
	GetBranchID() (*response.BranchResponse, error)
	GetInsuranceTypes() (*response.InsuranceTypesResponse, error)
	GetInsurance(params request.InsuranceRequest) (*response.InsuranceResponse, error)
	GetAssetCode(paramHeader request.HeaderAssetCodeRequest, reqBody request.AssetCodeRequest) (*response.AssetCodeResponse, error)
	GetProvinces(params request.ProvincesRequest) (*response.ProvincesResponse, error)
	GetCities(params request.CitiesRequest) (*response.CitiesResponse, error)
}

type dsfProgramService struct {
	dsfProgramRepo repo.DsfProgramRepo
}

func NewDsfProgramService(
	dsfProgramRepo repo.DsfProgramRepo,
) DsfProgramService {
	return &dsfProgramService{
		dsfProgramRepo: dsfProgramRepo,
	}
}

func (s *dsfProgramService) GetAdditionalInsurance() (*response.AdditionalInsuranceResponse, error) {

	result, err := s.dsfProgramRepo.GetAdditionalInsurance()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetPackageNames(params request.HeaderPackageNameRequest, queryParams request.ParamsPackageNameRequest) (*response.PackageNameResponse, error) {

	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetPackageNames(request.HeaderPackageNameRequest{
		ApplicationName: params.ApplicationName,
		AssetCode:       params.AssetCode,
		BranchCode:      params.BranchCode,
	}, request.ParamsPackageNameRequest{
		CarCondition: queryParams.CarCondition,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetCarConditions() (*response.CarConditionResponse, error) {

	result, err := s.dsfProgramRepo.GetCarConditions()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetPackages(paramHeader request.HeaderPackageRequest, reqBody request.PackageRequest) (*response.PackageResponse, error) {
	if err := paramHeader.Validate(); err != nil {
		return nil, err
	}
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetPackages(request.HeaderPackageRequest{
		ApplicationName: paramHeader.ApplicationName,
	}, request.PackageRequest{
		Brand:        reqBody.Brand,
		Model:        reqBody.Model,
		Variant:      reqBody.Variant,
		Province:     reqBody.Province,
		City:         reqBody.City,
		PackageName:  reqBody.PackageName,
		CarCondition: reqBody.CarCondition,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetVariants(paramHeader request.HeaderUnitByModelsRequest) (*response.VariantsResponse, error) {
	if err := paramHeader.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetVariants(request.HeaderUnitByModelsRequest{
		ApplicationName: paramHeader.ApplicationName,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetPaymentTypes() (*response.PaymentTypesResponse, error) {

	result, err := s.dsfProgramRepo.GetPaymentTypes()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetModels(params request.ModelsRequest) (*response.ModelsResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	result, err := s.dsfProgramRepo.GetModels(request.ModelsRequest{
		Brand: params.Brand,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetBrands(params request.BrandsRequest) (*response.BrandsResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}
	result, err := s.dsfProgramRepo.GetBrands(request.BrandsRequest{
		Keyword: params.Keyword,
		Limit:   params.Limit,
		Offset:  params.Offset,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetVehicleCategory() (*response.VehicleCategory, error) {
	result, err := s.dsfProgramRepo.GetVehicleCategory()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetBranchID() (*response.BranchResponse, error) {

	result, err := s.dsfProgramRepo.GetBranchID()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetInsuranceTypes() (*response.InsuranceTypesResponse, error) {

	result, err := s.dsfProgramRepo.GetInsuranceTypes()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetInsurance(params request.InsuranceRequest) (*response.InsuranceResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetInsurance(request.InsuranceRequest{
		DsfBranchId:       params.DsfBranchId,
		VehicleCategory:   params.VehicleCategory,
		InsuranceTypeCode: params.InsuranceTypeCode,
		CarCondition:      params.CarCondition,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetAssetCode(paramHeader request.HeaderAssetCodeRequest, reqBody request.AssetCodeRequest) (*response.AssetCodeResponse, error) {
	if err := paramHeader.Validate(); err != nil {
		return nil, err
	}
	if err := reqBody.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetAssetCode(request.HeaderAssetCodeRequest{
		ApplicationName: paramHeader.ApplicationName,
	}, request.AssetCodeRequest{
		VariantName:      reqBody.VariantName,
		CarCondition:     reqBody.CarCondition,
		ManufacturedYear: reqBody.ManufacturedYear,
		ModelName:        reqBody.ModelName,
		BrandName:        reqBody.BrandName,
	})

	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *dsfProgramService) GetProvinces(params request.ProvincesRequest) (*response.ProvincesResponse, error) {

	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetProvinces(request.ProvincesRequest{
		Search: params.Search,
		Offset: params.Offset,
		Limit:  params.Limit,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *dsfProgramService) GetCities(params request.CitiesRequest) (*response.CitiesResponse, error) {

	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.dsfProgramRepo.GetCities(request.CitiesRequest{
		Search:       params.Search,
		ProvinceCode: params.ProvinceCode,
		Offset:       params.Offset,
		Limit:        params.Limit,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
