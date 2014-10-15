package models

import (
	"github.com/creativelikeadog/revmgo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const COLLECTION = "books"

type Book struct {
	Id    bson.ObjectId `bson:"_id,omitempty"`
	Title string        `bson:"Title"`
	Body  string        `bson:"Body"`
	Tags  []string      `bson:"Tags"`
}

func Collection(d *mgo.Database) *mgo.Collection {
	return d.C(COLLECTION)
}

func (b *Book) FindById(d *mgo.Database, id string) *Book {

	b := new(Book)

	if bson.IsObjectIdHex(id) {
		Id := bson.ObjectIdHex(id)
		Collection(d).FindId(Id).One(b)
	}

	return b
}

func FindByTitle(d *mgo.Database, Title string) *Book {
	b := new(Book)
	Collection(d).Find(bson.M{"Title": Title}).One(b)
	return b
}

func (b *Book) Save(d *mgo.Database) error {
	_, err := Collection(d).Upsert(bson.M{"_id": b.Id}, b)
	return err
}

func (b *Book) Delete(d *mgo.Database) error {
	return Collection(d).RemoveId(b.Id)
}

func (b *Book) String() string {
	return b.Id.Hex() + ": " + b.Title
}

func GetBook(n string) *Book {
	b := new(Book)
	switch n {
	case "MobyDick":
		b.Id = bson.ObjectIdHex("51e9ad9749a1b71843000001")
		b.Title = "Moby Dick"
		b.Body = "Queequeg was a native of Rokovoko, an island far away to the West and South. It is not down in any map; true places never are.\n\nWhen a new-hatched savage running wild about his native woodlands in a grass clout, followed by the nibbling goats, as if he were a green sapling; even then, in Queequeg's ambitious soul, lurked a strong desire to see something more of Christendom than a specimen whaler or two. His father was a High Chief, a King; his uncle a High Priest; and on the maternal side he boasted aunts who were the wives of unconquerable warriors. There was excellent blood in his veins&mdash;royal stuff; though sadly vitiated, I fear, by the cannibal propensity he nourished in his untutored youth.\n\nA Sag Harbor ship visited his father's bay, and Queequeg sought a passage to Christian lands. But the ship, having her full complement of seamen, spurned his suit; and not all the King his father's influence could prevail. But Queequeg vowed a vow. Alone in his canoe, he paddled off to a distant strait, which he knew the ship must pass through when she quitted the island. On one side was a coral reef; on the other a low tongue of land, covered with mangrove thickets that grew out into the water. Hiding his canoe, still afloat, among these thickets, with its prow seaward, he sat down in the stern, paddle low in hand; and when the ship was gliding by, like a flash he darted out; gained her side; with one backward dash of his foot capsized and sank his canoe; climbed up the chains; and throwing himself at full length upon the deck, grappled a ring-bolt there, and swore not to let it go, though hacked in pieces.\n\nIn vain the captain threatened to throw him overboard; suspended a cutlass over his naked wrists; Queequeg was the son of a King, and Queequeg budged not. Struck by his desperate dauntlessness, and his wild desire to visit Christendom, the captain at last relented, and told him he might make himself at home. But this fine young savage&mdash;this sea Prince of Wales, never saw the Captain's cabin. They put him down among the sailors, and made a whaleman of him. But like Czar Peter content to toil in the shipyards of foreign cities, Queequeg disdained no seeming ignominy, if thereby he might happily gain the power of enlightening his untutored countrymen. For at bottom&mdash;so he told me&mdash;he was actuated by a profound desire to learn among the Christians, the arts whereby to make his people still happier than they were; and more than that, still better than they were. But, alas! the practices of whalemen soon convinced him that even Christians could be both miserable and wicked; infinitely more so, than all his father's heathens. Arrived at last in old Sag Harbor; and seeing what the sailors did there; and then going on to Nantucket, and seeing how they spent their wages in that place also, poor Queequeg gave it up for lost. Thought he, it's a wicked world in all meridians; I'll die a pagan.\n\nAnd thus an old idolator at heart, he yet lived among these Christians, wore their clothes, and tried to talk their gibberish. Hence the queer ways about him, though now some time from home."
		b.Tags = []string{"Herman Melville", "Classics"}
	case "AroundWorld":
		b.Id = bson.ObjectIdHex("51e9ad9749a1b71843000002")
		b.Title = "Around the World in 80 Days"
		b.Body = "\"The owners are myself,\" replied the captain.  \"The vessel belongs to me.\"\n\n\"I will freight it for you.\"\n\n\"No.\"\n\n\"I will buy it of you.\"\n\n\"No.\"\n\nPhileas Fogg did not betray the least disappointment; but the situation was a grave one.  It was not at New York as at Hong Kong, nor with the captain of the Henrietta as with the captain of the Tankadere.  Up to this time money had smoothed away every obstacle.  Now money failed.\n\nStill, some means must be found to cross the Atlantic on a boat, unless by balloon&mdash;which would have been venturesome, besides not being capable of being put in practice.  It seemed that Phileas Fogg had an idea, for he said to the captain, \"Well, will you carry me to Bordeaux?\"\n\n\"No, not if you paid me two hundred dollars.\""
		b.Tags = []string{"Jules Verne", "Classics", "Contemporary", "Action", "Adventure", "Suspense", "Fantasy"}
	case "PrincessMars":
		b.Id = bson.ObjectIdHex("51e9ae1749a1b71843000004")
		b.Title = "A Princess of Mars"
		b.Body = "Tal Hajus arose, and I, half fearing, half anticipating his intentions, hurried to the winding runway which led to the floors below.  No one was near to intercept me, and I reached the main floor of the chamber unobserved, taking my station in the shadow of the same column that Tars Tarkas had but just deserted.  As I reached the floor Tal Hajus was speaking.\n\n\"Princess of Helium, I might wring a mighty ransom from your people would I but return you to them unharmed, but a thousand times rather would I watch that beautiful face writhe in the agony of torture; it shall be long drawn out, that I promise you; ten days of pleasure were all too short to show the love I harbor for your race.  The terrors of your death shall haunt the slumbers of the red men through all the ages to come; they will shudder in the shadows of the night as their fathers tell them of the awful vengeance of the green men; of the power and might and hate and cruelty of Tal Hajus.  But before the torture you shall be mine for one short hour, and word of that too shall go forth to Tardos Mors, Jeddak of Helium, your grandfather, that he may grovel upon the ground in the agony of his sorrow.  Tomorrow the torture will commence; tonight thou art Tal Hajus'; come!\"\n\nHe sprang down from the platform and grasped her roughly by the arm, but scarcely had he touched her than I leaped between them.  My short-sword, sharp and gleaming was in my right hand; I could have plunged it into his putrid heart before he realized that I was upon him; but as I raised my arm to strike I thought of Tars Tarkas, and, with all my rage, with all my hatred, I could not rob him of that sweet moment for which he had lived and hoped all these long, weary years, and so, instead, I swung my good right fist full upon the point of his jaw.  Without a sound he slipped to the floor as one dead.\n\nIn the same deathly silence I grasped Dejah Thoris by the hand, and motioning Sola to follow we sped noiselessly from the chamber and to the floor above.  Unseen we reached a rear window and with the straps and leather of my trappings I lowered, first Sola and then Dejah Thoris to the ground below.  Dropping lightly after them I drew them rapidly around the court in the shadows of the buildings, and thus we returned over the same course I had so recently followed from the distant boundary of the city.\n\nWe finally came upon my thoats in the courtyard where I had left them, and placing the trappings upon them we hastened through the building to the avenue beyond.  Mounting, Sola upon one beast, and Dejah Thoris behind me upon the other, we rode from the city of Thark through the hills to the south.\n\nInstead of circling back around the city to the northwest and toward the nearest waterway which lay so short a distance from us, we turned to the northeast and struck out upon the mossy waste across which, for two hundred dangerous and weary miles, lay another main artery leading to Helium."
		b.Tags = []string{"Edgar Rice Burroughs", "Adventure"}
	case "EarthsCore":
		b.Id = bson.ObjectIdHex("51e9ae4949a1b71843000005")
		b.Title = "At the Earth's Core"
		b.Body = "With no heavenly guide, it is little wonder that I became confused and lost in the labyrinthine maze of those mighty hills.  What, in reality, I did was to pass entirely through them and come out above the valley upon the farther side.  I know that I wandered for a long time, until tired and hungry I came upon a small cave in the face of the limestone formation which had taken the place of the granite farther back.\n\nThe cave which took my fancy lay halfway up the precipitous side of a lofty cliff.  The way to it was such that I knew no extremely formidable beast could frequent it, nor was it large enough to make a comfortable habitat for any but the smaller mammals or reptiles.  Yet it was with the utmost caution that I crawled within its dark interior.\n\nHere I found a rather large chamber, lighted by a narrow cleft in the rock above which let the sunlight filter in in sufficient quantities partially to dispel the utter darkness which I had expected.  The cave was entirely empty, nor were there any signs of its having been recently occupied.  The opening was comparatively small, so that after considerable effort I was able to lug up a bowlder from the valley below which entirely blocked it.\n\nThen I returned again to the valley for an armful of grasses and on this trip was fortunate enough to knock over an orthopi, the diminutive horse of Pellucidar, a little animal about the size of a fox terrier, which abounds in all parts of the inner world.  Thus, with food and bedding I returned to my lair, where after a meal of raw meat, to which I had now become quite accustomed, I dragged the bowlder before the entrance and curled myself upon a bed of grasses&mdash;a naked, primeval, cave man, as savagely primitive as my prehistoric progenitors.\n\nI awoke rested but hungry, and pushing the bowlder aside crawled out upon the little rocky shelf which was my front porch.  Before me spread a small but beautiful valley, through the center of which a clear and sparkling river wound its way down to an inland sea, the blue waters of which were just visible between the two mountain ranges which embraced this little paradise.  The sides of the opposite hills were green with verdure, for a great forest clothed them to the foot of the red and yellow and copper green of the towering crags which formed their summit.  The valley itself was carpeted with a luxuriant grass, while here and there patches of wild flowers made great splashes of vivid color against the prevailing green."
		b.Tags = []string{"Edgar Rice Burroughs", "Adventure", "Action", "Fantasy", "Science Fiction"}
	case "WarWorlds":
		b.Id = bson.ObjectIdHex("51e9af2749a1b71843000006")
		b.Title = "The War of the Worlds Book I"
		b.Body = "\"Did you see a man in the pit?\" I said; but he made no answer to that.  We became silent, and stood watching for a time side by side, deriving, I fancy, a certain comfort in one another's company.  Then I shifted my position to a little knoll that gave me the advantage of a yard or more of elevation and when I looked for him presently he was walking towards Woking.\n\nThe sunset faded to twilight before anything further happened.  The crowd far away on the left, towards Woking, seemed to grow, and I heard now a faint murmur from it.  The little knot of people towards Chobham dispersed.  There was scarcely an intimation of movement from the pit.\n\nIt was this, as much as anything, that gave people courage, and I suppose the new arrivals from Woking also helped to restore confidence.  At any rate, as the dusk came on a slow, intermittent movement upon the sand pits began, a movement that seemed to gather force as the stillness of the evening about the cylinder remained unbroken.  Vertical black figures in twos and threes would advance, stop, watch, and advance again, spreading out as they did so in a thin irregular crescent that promised to enclose the pit in its attenuated horns.  I, too, on my side began to move towards the pit.\n\nThen I saw some cabmen and others had walked boldly into the sand pits, and heard the clatter of hoofs and the gride of wheels.  I saw a lad trundling off the barrow of apples.  And then, within thirty yards of the pit, advancing from the direction of Horsell, I noted a little black knot of men, the foremost of whom was waving a white flag.\n\nThis was the Deputation.  There had been a hasty consultation, and since the Martians were evidently, in spite of their repulsive forms, intelligent creatures, it had been resolved to show them, by approaching them with signals, that we too were intelligent."
		b.Tags = []string{"H. G. Wells", "Science Fiction", "Classics"}
	}
	return b
}
