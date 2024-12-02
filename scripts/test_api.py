#!/usr/bin/env python3

import requests
import base64
import json

url = "http://localhost:8080/synthesize/google"

payload = {
    # "text" : "จิตใจอ่อนแอเพียงใดเมื่อต้องการลืม บางทีคุณอาจไม่ลืม บางทีคุณอาจจะโกหก มันเป็นเรื่องโกหกที่คุณบอกทุกคนรอบตัวคุณ หรือบางทีอาจจะเป็นเรื่องโกหกที่คุณบอกตัวเอง - โยฮัน ลีเบิร์ต"
    "text" : "ว่าจะใด๋"
}

resp = requests.post(url, json=payload)

print( resp.status_code )

respDict = json.loads(resp.content)
audioDataStr = respDict['audio_data']

# print(type(respDict['audio_data']))
# audioDataByte = bytes(audioDataStr, 'utf-8')
audioDataByte = base64.b64decode(respDict['audio_data'])


with open("output_eiei.mp3", 'wb') as f:
    f.write(audioDataByte)
