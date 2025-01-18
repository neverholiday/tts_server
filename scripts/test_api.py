#!/usr/bin/env python3

import requests
import base64
import json
import time


text = "สวัสดีค่ะ เราเคยรู้จักกันรึเปล่า"

urlList = [
    ("http://localhost:8080/synthesize/google", 
     "google.mp3",
     {
         "text": text
     }),

]

# // VoiceAlloy   SpeechVoice = "alloy"
# // VoiceEcho    SpeechVoice = "echo"
# // VoiceFable   SpeechVoice = "fable"
# // VoiceOnyx    SpeechVoice = "onyx"
# // VoiceNova    SpeechVoice = "nova"
# // VoiceShimmer SpeechVoice = "shimmer"


for url, fName, payload in urlList:
    
    startTime = time.time()
    resp = requests.post(url, json=payload)
    print( f'API time usage: {time.time()-startTime} second' )

    if resp.status_code != 200:
        print( f'return code not 200: actual: [{resp.status_code}] {resp.content!r}' )
        continue

    respDict = json.loads(resp.content)
    audioDataStr = respDict['audio_data']
    audioDataByte = base64.b64decode(respDict['audio_data'])

    print(f'Write audio file: {fName}')
    with open(fName, 'wb') as f:
        f.write(audioDataByte)
