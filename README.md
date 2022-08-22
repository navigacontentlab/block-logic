# Blocklogic for NavigaDoc blocks

## Description
This library implements how to describe block logic using conditions for selecting Naviga Docs with a specific blocks.

You have a document with the following links: 
```
{
  "links": [
    {
      "uuid": "c8a565b3-f042-4cc7-8c61-df48b435fe1b",
      "type": "x-im/channel",
      "title": "IOL",
      "rel": "mainchannel"
    }
  ]
}
```

When following the event- or contentlog, you need to get the document in json from CCA and select it if the channel above exists. This library makes it possible to configure a condition like this:

```
{
  "in": "links",
    "or": [
        {
            "uuid": "c8a565b3-f042-4cc7-8c61-df48b435fe1b",
            "type": "x-im/channel",
            "title": "IOL",
            "rel": "mainchannel"
        },
        {
            "uuid": "81dd5b26-c1df-437c-969d-49ac65f426b5",
            "type": "x-im/channel",
            "title": "INL",
            "rel": "channel"
        }
    ]
}
```
The condition struct with and or condition:
```

```

Example of a more complex condition: 
```
{
  "in": "links",
  "or": [
    {
      "rel": "creator",
      "and": [
        {
          "in": "links",
          "rel": "affiliation",
          "and": [
            {
              "in": "links",
              "or": [
                {
                  "uri": "imid://unit/A"
                },
                {
                  "uri": "imid://unit/B"
                }
              ]
            }
          ]
        }
      ]
    },
    {
      "rel": "shared-with",
      "and": [
        {
          "in": "links",
          "or": [
            {
              "uri": "imid://unit/A"
            },
            {
              "uri": "imid://unit/B"
            }
          ]
        }
      ]
    }
  ]
}
```
Document must have either unit A OR unit B also shared with unit A or B
