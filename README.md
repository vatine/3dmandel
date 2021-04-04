# 3D Mandelbot analogue

It all started with "there is no good 3D analogue of eth
Mandelbrot". At some level, this is understandable, because the
easiest way of seeing the 2D Mandelbrot iteration is as "squaring and
adding" and a 3D vector multiplied by itself is, well, nothing.

But, not really happy with this, I started thinking about what
"squaring" and "adding" actually meant, geometrically, in the complex
space. Squaring (well, all multiplications) are simply "rotate and
scale" and adding is simply translation.

With this in mind, I spent some more time thinking what would be
needed. Specifically, we'd need a 3D transformation matrix, based on
the "last step", followed by a translation (this being the point in 3D
space that we're interested to see if it belongs to the set or not).

## Translation

A translation is simple, just add the point/vector that is the result
of the last transformation to the point we are checking for membership
in the test.

## Rotation and scaling

This is a bit more complicated. First, we define taht we're using a right-hand coordinate system. Then, we generate a transform matrix based on the previous poiint, the new point, and teh last transform.

1. Align a transform matrix so its first ("X" axis) is the latest
   point.
2. With this done, we can then define that we want the transform to be
   "in the XY plane" of the last translation. This, basically, means we
   take the cross-product of the previous point computed and the latest
   point computed. This gies uz the "Z" axis of out transform.

   However, we also want the next transform to be as close as possibel
   to the existing transform, so we also compute the negative of said
   Z axis, then pick the closes of Z and -Z as our new Z axis.
3. We can then generate the remaining axis of our transform matrix as the
   cross-product of our X and Z axes.
4. To get the scaling right, we then forcibly rescale the Y and Z axes to be
   the same length as the X axis.
