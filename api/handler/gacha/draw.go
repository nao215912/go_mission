package gacha

import (
	"github.com/gin-gonic/gin"
	"go_mission/api/domain/object"
	"go_mission/api/handler/auth"
	"go_mission/api/handler/httperror"
	"math/rand"
	"net/http"
)

type DrawRequest struct {
	Times int `json:"times"`
}

var characterSet = []string{"Hera", "Ares", "Athena", "Hermes", "Demeter", "Dionysus", "Aphrodite", "Apollo", "Artemis", "Eros", "Pan", "Harmonia", "Hebe", "Nike", "Asclepius", "Hecate", "Bia", "Deimos", "Phobos", "Alecto", "Megaera", "Tisiphone", "Vesta", "Thanatos", "Rhode", "Nyx", "Enyo", "Atropos", "Clotho", "Lachesis", "Gaia", "Iris", "Eos", "Chronus", "Terminus", "Janus", "Sylvanus", "Odin", "Thor", "Loki", "Baldr", "Hodur", "Hermod", "Frey", "Freya", "Forseti", "Bragi", "Hel", "Frigga", "Tyr", "Njord", "Mimir", "Urd", "Skuld", "Verandi", "Volla", "Vidar", "Ratatoskr", "Ra", "Khepri", "Osiris", "Anubis", "Isis", "Horus", "Set", "Bast", "Bes", "Sobek", "Nephthys", "Hathor", "Sekhmet", "Thoth", "Serqet", "Nepri", "Hapi", "Taweret", "Khonsu", "Medjed", "Kebechet", "Anhur", "Pakhet", "Ha", "Heh", "Nu", "Dua", "Min", "Shai", "Ptah", "Pan-Ku", "Shennong", "Nu Wa", "Yu Huang", "Sun Wukong", "Hou Tou", "Guan Yu", "Tian Mu", "Lei Gong", "Shou Hsing", "Xi Wangmu", "Kui Xing", "Zhu Rong", "Nezha", "Yen-Lo Wang", "Zao Shen", "Choy Sun Yeh", "Chang'e", "Eight Immortals", "Izanagi", "Izanami", "Amaterasu", "Tsukuyomi", "Susanoo", "Bishamon", "Hotei", "Ebisu", "Daikokuten", "Banzaiten", "Fukurokuju", "Jurōjin", "Inari", "Kaminari", "Raijin", "Fujin", "Kagutsuchi", "Ame-No-Mi-Kumari", "Amatsu-Mikaboshi", "Kuzenbo", "The Dagda", "Balor", "Angus", "Anpao", "Cerridwen", "Cernunnos", "Gwynn", "Danu", "Leir", "Fuamnach", "Bodb Derg", "Lugh", "Nuada", "Morrigan", "Arawn", "Badb", "Brigit", "Boann", "Taranis", "Artio", "Epona", "Cairbre", "Carman", "Dub", "Dother", "Dian", "Brahman", "Vishnu", "Shiva", "Kali-Ma", "Ganesha", "Ratli", "Yama", "Indra", "Agni", "Parvati", "Lakshmi", "Maya", "Kartikeya", "Ahura Mazda", "Mithras", "Ahriman", "Ametertat", "Atar", "Asha", "Armaiti", "Kshathra Vairya", "Vohu Manah", "Haurvatat", "Zurvan", "Gitche Manitou", "Hotamintanio", "Haokah", "Malsumis", "Calumet", "Nokomis", "Nanabozho", "Iktomi", "Owayodata", "Tawa", "Tomazooma", "Dzeharlons", "Kaiti", "Azeban", "Jiibayaabooz", "Hodiak", "Narya", "Nanuq", "Negafok", "Kadlu", "Nelvanna", "Turoq", "Sedna", "Tekkeitsertok", "Aipalovik Agloolik", "Quetzalcoatl", "Tezcatlipoca", "Huitzilopochtli", "Tlaloc", "Coatlicue", "Xipe Totec", "Ozomatli", "Xochipilli", "Xochiquetzal", "Michtlantecuhtli", "Xolotl", "Tonatiuh", "Huehuecoyotl", "Itzamna", "Hunab Ku", "Ixchel", "Wayep", "Chaac", "Ah Muzen Cab", "Awilix", "Hun Batz", "Acat", "Kukulkan", "Ah Kin Xoc", "Ixcacao", "Acan", "Ek Chuah", "Xbalanque", "Hunaphu", "Ah-Puch", "Buluc Chabtan", "Mam", "Inti", "Chasca", "Catequil", "Kon", "Pachamama", "Viracocha", "Supay", "Mama Cocha", "Mama Quilla", "Urcuchillay", "Axomamma", "Kane Milohai", "Ku", "Lono", "Kanaloa", "Papa", "Ragni", "Namaka", "Whiro", "Pele", "Poliahu", "B'ngudja", "Ukanipo", "Kamohoalii", "Apukohai", "Lakakane", "Kuape", "Maui", "Altjira", "Baiame", "Julunggul", "Daramulum", "Mamaragan", "Gnowee", "Marmoo", "Narahdarn", "Wambeen", "Bathala", "Apo Laki", "Mayari", "Amihan", "Anitun", "Aswang", "Detinos", "Aman Sinaya", "Apo Angin", "Tengri", "Ulgen", "Vaat", "Kuara", "Ay", "Koyash", "Erlik", "Gesar", "Kyazaghan", "Umay", "Kayra", "Kurmez", "Murgen", "Kubai", "Ak Ana", "Ay-Ata", "Gun Ana", "Yel Ana", "Yel Ata", "Berkut", "Aisyt", "Su Ana", "Su Ata", "Od Ana", "Od Ata", "Yer Tanry", "Etuhen", "Jaiuk", "Alaz", "Baianai", "Adaghan", "Akbugha", "Shalyk", "Inehsit", "Uren", "Qovaq", "Zarlik", "Zada", "Ukulan", "Izih", "Chokqu", "Talai", "Svarog", "Veles", "Perun", "Stribog", "Svarozvich", "Milda", "Saule", "Marzanna", "Dazhbog", "Laima", "Lada", "Svantovit", "Chernobog", "Buluku", "Avlekete", "Ogun", "Shango", "Damballah", "Anansi", "Eschu", "Oshun", "Ghekre", "Lusa", "Mahu", "Legba", "Baron Samedi", "Annu", "Ba'al", "Eriskegal", "Inanna", "Martu", "Nergal", "Ningal", "Saja", "Ninhursag", "Tammuz", "Marduk", "Tiamat", "Apsu", "Ukko", "Äkräs", "Nyyrikki", "Tapio", "Tuoni", "Vammatar", "Ilmarinen", "Ahti", "Akka", "Ajatar", "Lempo", "Hwanin", "Munsin", "Nulgupjisin", "Cheuksin", "Eopsin", "Jowangsin", "Magu", "Sosamsin", "Teojusin", "Ungnyeo", "Jumong", "Sumyeong-Jangja", "Yondung Halmoni", "Allah", "Allat", "Aluzza", "Manat", "Hubal", "Wadd", "Suwa'", "Yaghuth", "Ya'uq", "Nasr", "Nuha", "Tupa", "Sume", "Guaraci", "Jaci", "Ceuci", "Akuanduba", "Iara", "Yeba Belo", "Jurupari", "Anhanga", "Yorixiriamori", "Yahweh", "Satan", "Saint Michael", "Saint Gabriel"}

type DrawResponse struct {
	Results []*object.Character
}

func (h *handler) Draw(c *gin.Context) {
	ctx := c.Request.Context()
	req := new(DrawRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		httperror.BadRequest(c, err)
		return
	}

	user, err := auth.UserOf(c)
	if err != nil {
		httperror.BadRequest(c, err)
	}

	characters, err := h.app.Dao.Character().Create(ctx, characters(req.Times))
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}

	_, err = h.app.Dao.UserCharacter().Create(ctx, userCharacters(user, characters, req))
	if err != nil {
		httperror.InternalServerError(c, err)
		return
	}

	c.JSON(http.StatusOK, &DrawResponse{Results: characters})
}

func userCharacters(user *object.User, characters []*object.Character, req *DrawRequest) []*object.UserCharacter {
	userCharacters := make([]*object.UserCharacter, 0, req.Times)
	for i := 0; i < req.Times; i++ {
		userCharacter := &object.UserCharacter{
			UserID:      user.ID,
			CharacterID: characters[i].ID,
		}
		userCharacters = append(userCharacters, userCharacter)
	}
	return userCharacters
}

func characters(n int) []*object.Character {
	res := make([]*object.Character, 0, n)
	for i := 0; i < n; i++ {
		character := &object.Character{
			Name: characterSet[rand.Intn(len(characterSet))],
		}
		res = append(res, character)
	}
	return res
}
