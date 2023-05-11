package model

type UserModel struct {
	UserRecords []*UserEntity
}

type UserEntity struct {
	UxuyId    string
	Name      string
	Address   string
	Avatar    string
	Did       string
	CreatedAt string
	MpcKey    string
}

func NewUserModel() (*UserModel, error) {
	return &UserModel{
		UserRecords: initUserData(),
	}, nil
}

func (m *UserModel) FindUser(uxuyId string) (*UserEntity, error) {
	var user *UserEntity
	var err error
	for _, record := range m.UserRecords {
		if record.UxuyId == uxuyId {
			user = record
			break
		}
	}

	if user == nil {
		user, err = m.FindUserByAddress(uxuyId)
		if err != nil {
			return nil, err
		}
	}

	if user == nil {
		return nil, ErrNotFound
	}

	return user, nil
}

func (m *UserModel) FindUserByDid(did string) (*UserEntity, error) {
	var user *UserEntity
	for _, record := range m.UserRecords {
		if record.Did == did {
			user = record
			break
		}
	}

	if user == nil {
		return nil, ErrNotFound
	}

	return user, nil
}

func (m *UserModel) FindUserByAddress(address string) (*UserEntity, error) {
	var user *UserEntity
	for _, record := range m.UserRecords {
		if record.Address == address {
			user = record
			break
		}
	}

	if user == nil {
		return nil, ErrNotFound
	}

	return user, nil
}

func (m *UserModel) FindUsers() ([]*UserEntity, error) {

	return m.UserRecords, nil
}

func initUserData() []*UserEntity {
	userRecords := make([]*UserEntity, 6)
	userRecords[0] = &UserEntity{
		UxuyId:    "6ZuzN4MOiLjEqIsQ",
		Name:      "user1",
		Avatar:    "https://thirdwx.qlogo.cn/mmopen/vi_32/ajNVdqHZLLAkZDDuorBaRk7N5eFLY9z4QEsjjLaIfH5ITaAqibkfibkTalYrhh7cnoAcvl29VeAnJMALZEUhx64Q/132",
		Did:       "user1.uxuy",
		CreatedAt: "2023-05-10 10:54:11",
		Address:   "0x976cd24bAdDfC6B3025B006aB0c2B7b781407305",
		MpcKey:    "j783$nlHrlAmRrGHLNLIkr*7uKTboiTZL7EyDdBCHtBfEJN2Rx!bexsR989Wp0wx",
	}
	userRecords[1] = &UserEntity{
		UxuyId:    "Sird0uHHlxjvpyBq",
		Name:      "user2",
		Avatar:    "https://thirdwx.qlogo.cn/mmopen/vi_32/ajNVdqHZLLAkZDDuorBaRk7N5eFLY9z4QEsjjLaIfH5ITaAqibkfibkTalYrhh7cnoAcvl29VeAnJMALZEUhx64Q/132",
		Did:       "user2.uxuy",
		CreatedAt: "2023-05-10 10:54:12",
		Address:   "0x0A7B82E37EdFc8A5Df428C0654706Be761769b5f",
		MpcKey:    "J&Qo$Y*C2EUlLdTazfWHUoL#3TShGuHQt08Ne23A0a&IxtOGrb@s%0I7CYyg!sY*",
	}
	userRecords[2] = &UserEntity{
		UxuyId:    "neKWkHFCuLpRJVlV",
		Name:      "user3",
		Avatar:    "https://thirdwx.qlogo.cn/mmopen/vi_32/ajNVdqHZLLAkZDDuorBaRk7N5eFLY9z4QEsjjLaIfH5ITaAqibkfibkTalYrhh7cnoAcvl29VeAnJMALZEUhx64Q/132",
		Did:       "user3.uxuy",
		CreatedAt: "2023-05-10 10:54:13",
		Address:   "0x503f916C4Ca9FE67C6B540b74B280c0aAE3eFa77",
		MpcKey:    "a^!jRM7VvC2Sfq0k05nb1*uy7$oHQ40tDe1@*Ha8WdTBiEx1FmvRIXkSVBOi0Fzv",
	}
	userRecords[3] = &UserEntity{
		UxuyId:    "SNvx8du7MpBeqOoO",
		Name:      "user4",
		Avatar:    "https://thirdwx.qlogo.cn/mmopen/vi_32/ajNVdqHZLLAkZDDuorBaRk7N5eFLY9z4QEsjjLaIfH5ITaAqibkfibkTalYrhh7cnoAcvl29VeAnJMALZEUhx64Q/132",
		Did:       "user4.uxuy",
		CreatedAt: "2023-05-10 10:54:14",
		Address:   "0x3ecFaE5AEc20D9d4411296667c2a728fDC974525",
		MpcKey:    "sAu%#!gf2LT3Z$!qFEdne1SSWvMt0sfsgdvpxPHE9E@QIaEGRrnsWVdLrRsZVEAZ",
	}
	userRecords[4] = &UserEntity{
		UxuyId:    "hsiyQi5eJhoBCzqt",
		Name:      "user5",
		Avatar:    "https://thirdwx.qlogo.cn/mmopen/vi_32/ajNVdqHZLLAkZDDuorBaRk7N5eFLY9z4QEsjjLaIfH5ITaAqibkfibkTalYrhh7cnoAcvl29VeAnJMALZEUhx64Q/132",
		Did:       "user5.uxuy",
		CreatedAt: "2023-05-10 10:54:15",
		Address:   "0x3876AbC451fb98b1a88152e75B034D24fe1aF9B3",
		MpcKey:    "0RmBv7MJB9WFHg@gMQP7%hhpch0Wcjcz22&W2mINRcbilvj*RQ43h^a#AJwV2arU",
	}
	userRecords[5] = &UserEntity{
		UxuyId:    "1PrsOADSL4OMWMWi",
		Name:      "user6",
		Avatar:    "https://thirdwx.qlogo.cn/mmopen/vi_32/ajNVdqHZLLAkZDDuorBaRk7N5eFLY9z4QEsjjLaIfH5ITaAqibkfibkTalYrhh7cnoAcvl29VeAnJMALZEUhx64Q/132",
		Did:       "user6.uxuy",
		CreatedAt: "2023-05-10 10:54:16",
		Address:   "0xD142F7b8033084AbF2791446d4920fF06aC09e06",
		MpcKey:    "nwtzkm8*NfHRkXvgVzU*g@dOqM1dwWe26F7#NZqYF%jfk5HmE!yMRtwFXD6Rt9XQ",
	}
	return userRecords
}
