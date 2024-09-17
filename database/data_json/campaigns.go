package datajson

//https://json-generator.com/#
// [
//   '{{repeat(1, 20)}}',
//   {
//     "id": '{{index()}}',
//     "floor":'{{floating(9, 39, 2)}}',
//     "ad": {
//       "id": '{{guid()}}',
//       "title": '{{company()}}',
//       "duration": '{{random(15,20,30)}}'
//     },
//     "creative": {
//       "id": '{{objectId()}}',
//       "mediaFile": "http://testcontent.eyevinn.technology/ads/stswe19-teaser-15sek.mp4",
//       "w": 1920,
//       "h": 1080
//     }
//   }
// ]

func CampaignsData() []byte {
	return []byte(`[
  {
    "id": 0,
    "floor": 25.17,
    "ad": {
      "id": "8371dca7-610e-4944-a7c5-c7eb2e0309ad",
      "title": "Nurplex",
      "duration": 30
    },
    "creative": {
      "id": "644230e2b8a81234b744a8a3",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/stswe19-teaser-15sek.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 1,
    "floor": 16.71,
    "ad": {
      "id": "523a7ccc-327a-4c60-abf5-f74cbcd3742d",
      "title": "Ebidco",
      "duration": 20
    },
    "creative": {
      "id": "644230e2b982395004b3cf33",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/apotea-15s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 2,
    "floor": 35.79,
    "ad": {
      "id": "5d777e0f-8340-4e59-b26a-3690dc57f1ba",
      "title": "Dancity",
      "duration": 20
    },
    "creative": {
      "id": "644230e2b157c3f8d9a4d926",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/grannyra-10s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 3,
    "floor": 27.46,
    "ad": {
      "id": "67bace8a-e690-4eed-9ca3-7245050bd4cf",
      "title": "Exozent",
      "duration": 30
    },
    "creative": {
      "id": "644230e23ed7a300c7be7759",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/sff-15s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 4,
    "floor": 13.63,
    "ad": {
      "id": "abfe91cf-ccfe-4ad7-91ff-7ebf22d109e2",
      "title": "Zyple",
      "duration": 20
    },
    "creative": {
      "id": "644230e2b673920e79734c8a",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/alvedon-10s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 5,
    "floor": 21.1,
    "ad": {
      "id": "1676a5d5-efcf-4644-8558-5b3af7454747",
      "title": "Cowtown",
      "duration": 15
    },
    "creative": {
      "id": "644230e21e49f190cfb5521e",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/bromwel-15s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 6,
    "floor": 23.77,
    "ad": {
      "id": "77b47b4c-ed91-4a25-a2f3-e9c1d5977d68",
      "title": "Rugstars",
      "duration": 20
    },
    "creative": {
      "id": "644230e2331751780822638f",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/mio-15s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 7,
    "floor": 37.12,
    "ad": {
      "id": "374186ee-5efe-4a7e-b612-dd5e9059ec52",
      "title": "Schoolio",
      "duration": 30
    },
    "creative": {
      "id": "644230e2b1f930ab4ff1293b",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/volvo-15s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 8,
    "floor": 29.73,
    "ad": {
      "id": "f6cdf80d-0381-430d-9a3e-2c661199c644",
      "title": "Entropix",
      "duration": 15
    },
    "creative": {
      "id": "644230e2daec640873f7dc05",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/willys-20s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 9,
    "floor": 29.33,
    "ad": {
      "id": "ea3c446e-1570-4fdb-b17d-b170946e10b6",
      "title": "Pasturia",
      "duration": 20
    },
    "creative": {
      "id": "644230e26dd83fc1425c8cc8",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/sector-15s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 10,
    "floor": 11.61,
    "ad": {
      "id": "6c1fa370-24ed-46ed-8861-70c2c54ddbdb",
      "title": "Qualitern",
      "duration": 30
    },
    "creative": {
      "id": "644230e295983b5afc8e4046",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/kungsangen-10s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 11,
    "floor": 25.27,
    "ad": {
      "id": "c93fe204-b2cd-407c-862a-620b16db436e",
      "title": "Radiantix",
      "duration": 20
    },
    "creative": {
      "id": "644230e21a5f875e822475cd",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/nipaxon-20s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 12,
    "floor": 18.95,
    "ad": {
      "id": "4bb6be63-3377-4a97-bfb3-9366a19f72f4",
      "title": "Deepends",
      "duration": 20
    },
    "creative": {
      "id": "644230e28804a8ad8f2a5605",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/specsavers-15s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 13,
    "floor": 13.13,
    "ad": {
      "id": "681b86ff-23e6-4e9e-b0c8-2092df7569e9",
      "title": "Corpulse",
      "duration": 20
    },
    "creative": {
      "id": "644230e2cc4da9b609a60077",
      "mediaFile": "http://testcontent.eyevinn.technology/ads/coldzyme-10s.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 14,
    "floor": 25.55,
    "ad": {
      "id": "72a566fa-d20b-4407-a72b-26263f5bfc94",
      "title": "Spacewax",
      "duration": 30
    },
    "creative": {
      "id": "644230e2b99055a745c98409",
      "mediaFile": "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerEscapes.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 15,
    "floor": 21.25,
    "ad": {
      "id": "b1e58642-7fca-4b1d-b0fa-c664fe0369ed",
      "title": "Slambda",
      "duration": 20
    },
    "creative": {
      "id": "644230e288d5f2598886012a",
      "mediaFile": "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerJoyrides.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 16,
    "floor": 20.75,
    "ad": {
      "id": "e0e1fc19-75eb-4e63-b1b5-3126b591ace9",
      "title": "Proxsoft",
      "duration": 20
    },
    "creative": {
      "id": "644230e2c1bf4d5ec45cd828",
      "mediaFile": "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerMeltdowns.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 17,
    "floor": 12.91,
    "ad": {
      "id": "53e61312-75b9-4edf-881e-76b7df4589e8",
      "title": "Peticular",
      "duration": 20
    },
    "creative": {
      "id": "644230e210a32248738b130c",
      "mediaFile": "http://commondatastorage.googleapis.com/gtv-videos-bucket/sample/ForBiggerBlazes.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 18,
    "floor": 16.82,
    "ad": {
      "id": "0bba97fe-9174-4cf8-a931-3f2120a27819",
      "title": "Zerology",
      "duration": 15
    },
    "creative": {
      "id": "644230e23d9b62f9ac82e29a",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/NY2njqq_1625838402813.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 19,
    "floor": 27.76,
    "ad": {
      "id": "ebf28f3f-2dfc-46e6-908d-8ed63395686f",
      "title": "Krog",
      "duration": 15
    },
    "creative": {
      "id": "644230e23ef7bd4849115043",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/Mjfg9jB_1625838400106.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 20,
    "floor": 25.67,
    "ad": {
      "id": "d9f85ae5-da42-4360-8305-6a73e5fd64e5",
      "title": "Temorak",
      "duration": 15
    },
    "creative": {
      "id": "644230e207929179f95f2399",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/DZBXLCR_1625838397194.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 21,
    "floor": 20.66,
    "ad": {
      "id": "d8b36cc0-d5ab-4938-bb34-a618114b2dc5",
      "title": "Zillatide",
      "duration": 30
    },
    "creative": {
      "id": "644230e2afe1eb8e6a982ed9",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/REanQ1Q_1625838394125.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 22,
    "floor": 29.47,
    "ad": {
      "id": "5b15822b-5961-454e-9719-82c49a2ac63a",
      "title": "Kongene",
      "duration": 20
    },
    "creative": {
      "id": "644230e28eac69fd308c338e",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/pvgZo16_1625838391553.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 23,
    "floor": 33.52,
    "ad": {
      "id": "29591dc9-1d3e-4637-9b92-c85e3af4c67c",
      "title": "Calcu",
      "duration": 30
    },
    "creative": {
      "id": "644230e2d8e2de03643348e4",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/SCR8mqz_1625838388211.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 24,
    "floor": 17.23,
    "ad": {
      "id": "8290a471-8561-49c4-af7f-4424462f3644",
      "title": "Zolar",
      "duration": 30
    },
    "creative": {
      "id": "644230e2548e2a11a64ec80f",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/OztROWo_1625838384350.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 25,
    "floor": 33.64,
    "ad": {
      "id": "2ce5f4af-34b1-43f0-a406-ceff76782368",
      "title": "Rockabye",
      "duration": 30
    },
    "creative": {
      "id": "644230e26e76222b81e6225a",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/UZzAtO6_1625838381505.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 26,
    "floor": 14.89,
    "ad": {
      "id": "e9de1c0e-2c62-47d5-8fc0-1bc2acba4dfa",
      "title": "Zentime",
      "duration": 30
    },
    "creative": {
      "id": "644230e2fe548cd953c4f46a",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/U2hjjrA_1625838377830.mp4",
      "w": 1920,
      "h": 1080
    }
  },
  {
    "id": 27,
    "floor": 23.78,
    "ad": {
      "id": "fa90ae65-ce5e-4be5-abee-cf793ccf60b0",
      "title": "Aquacine",
      "duration": 20
    },
    "creative": {
      "id": "644230e2171b42b7246da453",
      "mediaFile": "https://cdn-cfy-p0.iqm.com/Video/1/620/Em6iTw3_1625838370039.mp4",
      "w": 1920,
      "h": 1080
    }
  }
]`)
}
