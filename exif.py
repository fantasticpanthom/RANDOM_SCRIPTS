import os
from PIL import Image
from PIL.ExifTags import TAGS
imagepath = input("Enter the path to the image")
if (os.path.isfile(imagepath)):
    for (i,j) in Image.open(imagepath)._getexif().items():
        print('%s = %s' % (TAGS.get(i), j))
else:
    print("IMAGE NOT FOUND.")        